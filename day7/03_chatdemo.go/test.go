package main

import (
	"fmt"
	"os"
	// "strings"
)

func main() {
	//定义缓冲区1024,用于接收键盘的输入内容
	buf := make([]byte, 1024)
	in := os.Stdin
	n, err := in.Read(buf)
	if err != nil {
		fmt.Println("read from stdin err :", err)
		return
	}
	fmt.Println(string(buf[:n]))
}
