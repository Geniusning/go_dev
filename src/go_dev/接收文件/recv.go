package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

//RecvFile 接收发送的文件内容
func RecvFile(path string, conn net.Conn) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("os.Create err=", err)
		return
	}
	defer file.Close()

	//读多少，写多少
	buf := make([]byte, 1024*4)
	for {
		n, err1 := conn.Read(buf)
		if err1 != nil {
			if err1 == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("os.Create err1=", err1)
			}
			return
		}

		if n == 0 {
			fmt.Println("接受完毕")
			break
		}
		file.Write(buf[:n])
	}

}
func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listener.Close()

	//阻塞等待用户发送内容
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("listener.Accept err1=", err1)
		return
	}
	defer conn.Close()
	//接收发送的名字
	buf := make([]byte, 1024)
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("listener.Accept err2=", err2)
		return
	}
	filename := string(buf[:n])
	//发送ok
	conn.Write([]byte("ok"))

	//接收发送的文件
	RecvFile(filename, conn)
}
