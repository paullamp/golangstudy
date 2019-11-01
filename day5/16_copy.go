package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 获取命令行输入
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s srcfile dstfile\n", os.Args[0])
		return
	}
	//打开源文件
	srcfile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("open src file failed", err)
		return
	}

	//关闭源文件
	defer srcfile.Close()
	//打开目标文件
	dstfile, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println("create dst file failed", err)
		return
	}
	//关闭目标文件
	defer dstfile.Close()

	//创建缓存字符切片，用来接收文件
	buf := make([]byte, 4*1024) // 定义4k的缓冲区，用于接收读写的文件
	var n int

	for {
		n, err = srcfile.Read(buf)
		if err == io.EOF {
			return
		} else if err != nil {
			fmt.Println("Read file err : ", err)
			return
		}
		dstfile.Write(buf[:n])
	}
}
