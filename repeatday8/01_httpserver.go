package main

import (
	"fmt"
	"net"
)

func main() {
	//启动监听
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("net listen failed")
	}

	//启动等待接收客户的请求
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("Listen and Accept faield")
		return
	}

	// 准备缓冲空间，以存储数据
	buf := make([]byte, 1024)
	n, errRead := conn.Read(buf)
	if errRead != nil {
		fmt.Println("Read failed")
		return
	}
	fmt.Println(string(buf[:n]))
}
