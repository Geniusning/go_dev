package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func sendFile(path string, conn net.Conn) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open files filed,err=", err)
		return
	}
	defer file.Close()

	//读文件,读多少，发送多少
	buf := make([]byte, 1024*4)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("发送完毕")
			} else {
				fmt.Println("read err=", err)
			}
			return
		}
		conn.Write(buf[:n])
	}
}
func main() {
	//输入文件名
	fmt.Println("请输入文件路径：")
	var path string
	fmt.Scan(&path)

	//获取文件名
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("fileInfo err= ", err)
		return
	}

	//连接服务器
	conn, err1 := net.Dial("tcp", "127.0.0.1:8000")
	if err1 != nil {
		fmt.Println("net.Dial err=", err1)
		return
	}
	defer conn.Close()

	//给接收方发送文件名，如果接收到ok,则开始发送文件
	_, err = conn.Write([]byte(fileInfo.Name()))
	if err != nil {
		fmt.Println("write err= ", err)
		return
	}

	//收到接收方回复
	var n int
	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn read err= ", err)
		return
	}
	//收到ok后发送文件数据
	fmt.Println("收到服务器的", string(buf[:n]))
	if "ok" == string(buf[:n]) {
		sendFile(path, conn)
	}
}
