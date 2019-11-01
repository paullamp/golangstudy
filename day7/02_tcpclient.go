package main

import (
	"fmt"
	"net"
)

func main() {
	//打开一个到服务器端的连接
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("connect to server error: ", err)
		return
	}
	defer conn.Close()
	conn.Write([]byte("Hello china"))
}
