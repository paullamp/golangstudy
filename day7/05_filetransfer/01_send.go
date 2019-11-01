package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

//将文件内容从本地读取，　并且发送到服务器
func sendfile(path string, conn net.Conn) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open error: ", err)
		return
	}
	defer f.Close()
	//循环读取文件并且发送
	buf := make([]byte, 1024*4)
	for {
		n, f_read_err := f.Read(buf)
		if f_read_err != nil {
			if f_read_err == io.EOF {
				fmt.Println("文件读取完毕")
				break
			} else {
				fmt.Println("f_read_err :", f_read_err)
				return
			}
		}
		conn.Write(buf[:n])
	}
}

func main() {
	// 1. 从命令行获取输入参数
	var path string
	fmt.Println("请输入需要传送的文件名：")
	fmt.Scan(&path)

	// 2. 从输入获取文件信息，并判断文件信息
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat error: ", err)
		return
	}

	//3. 连至服务器
	conn, conn_err := net.Dial("tcp", "127.0.0.1:8000")
	defer conn.Close()
	if conn_err != nil {
		fmt.Println("net.Dial error: ", conn_err)
		return
	}
	//3. 传输文件名称至服务器
	_, write_err := conn.Write([]byte(info.Name()))
	if write_err != nil {
		fmt.Println("conn.Write error:", write_err)
		return
	}
	//4. 获取服务器端返回的"ok"字符
	buf := make([]byte, 1024)
	n, read_err := conn.Read(buf)
	if read_err != nil {
		fmt.Println("conn.Read error: ", read_err)
		return
	}
	if "ok" == string(buf[:n]) {
		fmt.Println("服务器已准备好接收")
		sendfile(path, conn)
	}
}
