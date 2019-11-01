package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string // 用于发送信息的管道
	Name string      // 用户名
	Addr string      // 客户端的地址
}

var onlineMap map[string]Client

var message = make(chan string)

//转发消息，只要有消息来了，遍历每个map成员，发送消息
func Manager() {
	//给map分配空间
	onlineMap = make(map[string]Client)
	for {
		msg := <-message //没有消息前，这里会阻塞

		//遍历每个map成员，发送消息
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

//给客户端发送信息
func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ": " + msg
	return
}

//处理连接请求
func HandleConn(conn net.Conn) {
	//获取客户端地址
	cliaddr := conn.RemoteAddr().String()
	//创建一个客户端结构体变量,用于存储客户端的信息,初始未定义情况下，用户名和网络地址一样
	cli := Client{make(chan string), cliaddr, cliaddr}
	//添加到结构中
	onlineMap[cliaddr] = cli
	//给当前客户端发送信息
	go WriteMsgToClient(cli, conn)
	//广播某用户上线消息
	// message <- "[" + cli.Addr + "]" + cli.Name + ": login"
	message <- MakeMsg(cli, "login")
	cli.C <- MakeMsg(cli, "I am here")
	//新开一个协程，接收用户发送过来的数据
	isQuit := make(chan bool) // 判断用户是否是主动退出
	hasData := make(chan bool)
	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if n == 0 { //不管客户端断开或是出问题
				isQuit <- true
				fmt.Println("conn.read error = ", err)
				return

			}
			msg := string(buf[:n-1]) //windows 后面会多一个换行
			if len(msg) == 3 && msg == "who" {
				conn.Write([]byte("user list\n"))
				for _, clivalue := range onlineMap {
					// cli.C <- clivalue.Name
					msg := clivalue.Addr + ":" + clivalue.Name + "\n"
					conn.Write([]byte(msg))

				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				//重命名的格式rename |mike
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				onlineMap[cliaddr] = cli // 需要想一下，为什么还需要再赋一次值
				conn.Write([]byte("name changed to :" + name))

			} else {
				message <- MakeMsg(cli, msg)
			}
			hasData <- true //代表有数据
		}
	}()
	//死循环，保证此协程不退出
	for {
		//通过select 检查isquit ，　同时检测超时
		select {
		case <-isQuit:
			delete(onlineMap, cliaddr)
			message <- MakeMsg(cli, cliaddr+" quit") //广播谁下线了
		case <-hasData:

		case <-time.After(30 * time.Second):
			delete(onlineMap, cliaddr)
			message <- MakeMsg(cli, cliaddr+" timeout ") //广播超时下线
			return
		}
	}
}

//主进程
func main() {
	// 建立监听
	listener, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		fmt.Println("net.Listen error = ", err)
		return
	}
	defer listener.Close()
	go Manager()
	//循环阻塞等待接收连接
	for {
		conn, conn_err := listener.Accept()
		if conn_err != nil {
			fmt.Println("listener.Accept error = ", conn_err)
		}
		go HandleConn(conn)
	}
}
