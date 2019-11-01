package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}
	defer ln.Close()
	conn, err1 := ln.Accept()
	if err1 != nil {
		fmt.Println("ln.Accept err = ", err)
		return
	}
	defer conn.Close()
	buf := make([]byte, 1024*4)
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("read from browser error = ", err2)
		return
	}
	fmt.Printf("request mesage is : #%v#", string(buf[:n]))
}
