package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, "connected successfully")

	//处理数据返回大写字母
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn read err = ", err)
			return
		}
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main() {
	//监听客户端连接
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("lister err =", err)
		return
	}
	//关闭监听
	defer listener.Close()
	//接收客户端数据。新建
	for {
		conn, err1 := listener.Accept()
		if err1 != nil {
			fmt.Println("listener.Accept err = ", err)
			return
		}
		go HandleConn(conn)
	}
}
