package main

import (
	"fmt"
	"net"
	"time"
)

func tcpServer() {
	serveraddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("socket resolve failed")
	}
	ln, err := net.ListenTCP("tcp", serveraddr)
	if err != nil {
		fmt.Println("listen failed")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}

func tcpClient() {
	serveraddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, err := net.DialTCP("tcp", nil, serveraddr)
	if err != nil {
		fmt.Println("client ", err.Error())
	}
	b := make([]byte, 100)
	conn.Read(b)
	fmt.Println(b)
}

func main() {
	tcpServer()
}
