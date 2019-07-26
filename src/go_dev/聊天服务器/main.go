package main

import (
	"fmt"
	"net"
)

//Client is ClientInfo
type Client struct {
	C    chan string
	Name string
	Addr string
}

var onlineMap map[string]Client

var message chan string

//MakeMsg 处理消息
func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ":" + msg
	return
}
func NotifyOnline() {
	onlineMap = make(map[string]Client)
	for {
		msg := <-message //没有消息钱阻塞
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}
func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

//HandleConn 处理用户连接
func HandleConn(conn net.Conn) {
	defer conn.Close()
	ClientAddr := conn.RemoteAddr().String()

	//初始化用户
	cli := Client{make(chan string), ClientAddr, ClientAddr}
	onlineMap[ClientAddr] = cli

	//开一个携程，专门给客户端转发消息
	go WriteMsgToClient(cli, conn)
	//广播某个在线
	message <- MakeMsg(cli, "login")

	//新建一个携程，接收用户发送的消息
	go func() {
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println(" conn.Read err = ", err)
			return
		}

		msg := string(buf[:n-1])

		message <- MakeMsg(cli, msg)
	}()
	for {
	}

}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	defer listener.Close()

	// 遍历onlineMap,通知每一个人
	go NotifyOnline()
	//循环阻塞用户连接
	for {
		conn, err1 := listener.Accept()
		if err1 != nil {
			fmt.Println("listener.Accept err1 = ", err1)
			return
		}
		go HandleConn(conn)
	}

}
