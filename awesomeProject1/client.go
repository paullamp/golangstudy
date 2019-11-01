package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("ipv4", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("connect error")
		return
	}
	buf := make([]byte, 1024)
	n, connReadError := conn.Read(buf)
	if connReadError != nil {
		fmt.Println("connReadError error")
		return
	}
	fmt.Println("Server: ", string(buf[:n]))

}
