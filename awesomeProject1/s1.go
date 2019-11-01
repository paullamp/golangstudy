package main

import (
	"fmt"
	"net"
)

func main() {
	//建立本地监听服务
	ln, lnerr := net.Listen("tcp", "127.0.0.1:888")
	defer ln.Close()
	if lnerr != nil {
		fmt.Println("Listen error : ", lnerr)
		return
	}
	//通过accept 等待客户端连接
	conn, connerr := ln.Accept()
	defer conn.Close()
	if connerr != nil {
		fmt.Println("Conn error: ", connerr)
		return
	}
	//往客户端发送信息
	conn.Write([]byte("Server: This is message from server"))
	//接收客户端信息
	//创建1K缓存区，以接收客户端数据
	buf := make([]byte, 1024)
	n, connReadErr := conn.Read(buf)
	if connReadErr != nil {
		fmt.Println("connReadErr: ", connReadErr)
		return
	}
	fmt.Println(string(buf[:n]))

}
