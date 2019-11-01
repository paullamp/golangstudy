package main

import (
	"fmt"
	"net"
	"os"
	// "time"
)

func main() {
	serveraddr := "127.0.0.1:8000"
	conn, err := net.Dial("tcp", serveraddr)
	if err != nil {
		fmt.Println("conn err:", err)
		return
	}
	defer conn.Close()
	//创建一个协程，从服务器端获取输入
	messagefromserver := make([]byte, 1024*4)
	go func() {
		for {
			n, err := conn.Read(messagefromserver)
			if err != nil {
				fmt.Println("conn read error ", err)
				return
			}
			fmt.Println("Message from server is : ", string(messagefromserver[:n]))
		}
	}()

	// 从标准输入获取输入内容，并且写入到conn中,相当于在主进程中
	buf := make([]byte, 1024*4)
	for {
		inputn, inputerr := os.Stdin.Read(buf)
		if inputerr != nil {
			fmt.Println("error in read from stdin ", inputerr)
			return
		}
		conn.Write(buf[:inputn])

	}
}
