package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//判断输入参数
	if len(os.Args) != 2 {
		fmt.Println("输入的参数有误")
		fmt.Println("命令格式为：", os.Args[0], "[45|46]")
		return
	}
	node := os.Args[1]
	if node != "45" && node != "46" {
		fmt.Println("输入的待切换节点有误")
		fmt.Println("命令格式为：", os.Args[0], "[45|46]")
	}
	//连接服务器10.0.197.18:33445
	serverip := "10.0.197.18:33445"
	conn, errDial := net.Dial("tcp", serverip)
	if errDial != nil {
		fmt.Println("net.Dial to Server error:", errDial)
		return
	}
	defer conn.Close()
	command := []byte("/root/bash/bpm_switch_node.sh " + node)
	_, errWrite := conn.Write(command)
	if errWrite != nil {
		fmt.Println("conn.Write to server error:", errWrite)
		return
	}
	result := make([]byte, 1024)
	n, errRead := conn.Read(result)
	if errRead != nil {
		fmt.Println("服务器端返回错误:", errRead)
		return
	}
	fmt.Println(string(result[:n]))

}
