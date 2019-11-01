package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

//从流中读取内容并写入到文件中
func recvFile(filename string, conn net.Conn) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("create file error ", err)
		return
	}
	defer file.Close()
	buf := make([]byte, 1024*4)
	for {
		n, connreaderr := conn.Read(buf)
		if connreaderr != nil {
			if connreaderr == io.EOF {
				fmt.Println("File recv finished\n")
				return
			}
			fmt.Println("recvFile conn.Read error ", connreaderr)
			return
		}
		file.Write(buf[:n])
	}
}

func main() {
	//监听服务
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.listen error ", err)
		return
	}
	defer listener.Close()
	//接收请求
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("listener.accept error ", err1)
		return
	}
	defer conn.Close()
	//接收文件内容
	buf := make([]byte, 1024)
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("conn.Read error ", err2)
		return
	}

	filename := string(buf[:n])

	//发送一个ok
	_, err3 := conn.Write([]byte("ok"))
	if err3 != nil {
		fmt.Println("write ok message error ", err3)
		return
	}
	//接收文件
	recvFile(filename, conn)
}
