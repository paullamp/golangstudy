package main

import (
	"fmt"
	"net"
)

func main(){
	fmt.Println("welcome to golang web")
	ln,err := net.Listen("tcp","127.0.0.1:8000")
	if err != nil{
		fmt.Println("error found",err)
		return
	}
	conn,connErr := ln.Accept()
	if connErr != nil{
		fmt.Println("conn_error=",connErr)
		return
	}
	fmt.Println(conn)
	conn.Write([]byte("Hello world , client"))
	buf := make([]byte, 1024 *4)
	n,readErr := conn.Read(buf)
	if readErr != nil {
		return
	}
	fmt.Println("conn_read",  string(buf[:n]))


}
