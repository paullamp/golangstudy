package main

import (
	"log"
	"net"
	"os"
	"os/exec"
	"strings"

	_ "github.com/icattlecoder/godaemon"
)

const logfile = "/var/log/golangCommandServer.log"

func HandleCmd(conn net.Conn) {
	defer conn.Close()
	//接收客户端的请求，输出需要的命令
	//保存客户端输入的命令：
	commands := make([]byte, 1024)
	n, errRead := conn.Read(commands)
	if errRead != nil {
		log.Println("Read from client error:", errRead)
		return
	}
	//打印客户端输入的命令
	log.Println(string(commands[:n]))
	//调用本地脚本，获取输出信息，并且返回给客户端
	coms := string(commands[:n])
	comsSlice := strings.Split(coms, " ")

	cmd := exec.Command("bash", comsSlice[0], comsSlice[1])
	cmdresult, errCmdOutput := cmd.Output()
	if errCmdOutput != nil {
		log.Println("命令执行失败:", errCmdOutput)
		return
	}
	conn.Write(cmdresult)

	log.Println(string(cmdresult))
}

//　监听请求，等待用户输入，将用户输入转换成bash 命令
func main() {
	//打开日志文件，并且将日志的输出指定到文件
	logfilehandle, errOsCreate := os.Create(logfile)
	if errOsCreate != nil {
		log.Println("Open file error:", errOsCreate)
		return
	}
	defer logfilehandle.Close()
	log.SetOutput(logfilehandle)
	log.Println("TestLog")
	ln, err := net.Listen("tcp", ":33445")
	if err != nil {
		log.Println("net.Listen Error:", err)
		return
	}
	defer ln.Close()
	for {
		conn, errAccept := ln.Accept()
		if errAccept != nil {
			log.Println("ln.Accept Error:", err)
			return
		}
		go HandleCmd(conn)
	}

}
