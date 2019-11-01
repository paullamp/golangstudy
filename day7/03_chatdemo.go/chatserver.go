package main

import (
	"fmt"
	"net"
	"strings"
)

func handleclient(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, " connected")
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read err", err)
			return
		}
		// 判断输入是否为exit
		if strings.TrimSpace(string(buf[:n])) != "exit" {
			fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))
		} else {
			return
		}

		// conn.Write(buf[:n])
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))

	}
}

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("listen err: ", err)
		return
	}
	defer ln.Close()

	// 接收请求
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("conn error:", err)
			continue
		}
		go handleclient(conn)
	}

}
