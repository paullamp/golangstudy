package main

import (
	"fmt"
	"net"
)

func main() {
	ln, lnerr := net.Listen("tcp",":8899")
	if lnerr != nil {
		fmt.Println("Readerror", lnerr)
		return
	}
	defer ln.Close()
	conn, connerr := ln.Accept()
	if connerr != nil {
		fmt.Println(connerr)
		return
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	n, readerr := conn.Read(buf)
	if readerr != nil {
		fmt.Println("readerr:", readerr)
		return
	}
	fmt.Println(string(buf[:n]))
}
