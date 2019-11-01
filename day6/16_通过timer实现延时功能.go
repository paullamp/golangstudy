package main

import (
	"fmt"
	"time"
)

func main() {
	<-time.After(time.Second * 2) // 直接返回一个只读的channel , 在定时到后往channel中写数据
	fmt.Println("time is up")
}
func main02() {
	//直接通过sleep函数
	time.Sleep(time.Second * 2)
	fmt.Println("time is up")
}

//通过timer 实现定时功能，类似于time.sleep，等待什么时间后执行
func main01() {
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("time is up")
}
