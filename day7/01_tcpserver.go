package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8000") //申请监听在127.0.0.1:8000端口上
	defer ln.Close()
	if err != nil {
		fmt.Println("Listen err ", err)
		return
	}
	//阻塞等待接收用户请求
	conn, conn_err := ln.Accept()
	defer conn.Close()
	if conn_err != nil {
		fmt.Println("Conn error: ", conn_err)
		return
	}
	buf := make([]byte, 1024) //创建1k缓冲区大小，用于接收内容
	n, readerr := conn.Read(buf)
	if readerr != nil {
		fmt.Println("Read error: ", readerr)
		return
	}
	fmt.Println(string(buf[:n]))
}
