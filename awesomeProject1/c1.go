package main

import (
	"fmt"
	"net"
)

func main() {
	//连接远程server
	conn, err := net.Dial("tcp", "127.0.0.1:888")
	defer conn.Close()
	if err != nil {
		fmt.Println("conn err: ", err)
		return
	}
	//读取conn的内容
	buf := make([]byte, 1024)
	n, connReadErr := conn.Read(buf)
	if connReadErr != nil {
		fmt.Println("conn Read Err :", connReadErr)
		return
	}
	fmt.Println(string(buf[:n]))
	//向conn中写入内容
	conn.Write([]byte("What is your name?"))

}
