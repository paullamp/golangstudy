package main

import (
	"fmt"
	"time"
)

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Println()
}

var c chan int

// c := make(chan int) // 会报错，在函数外无法使用自动推导类型
func person1() {
	Printer("helloworld")
	c <- 99
}
func person2() {
	<-c
	Printer("learngolang")
}
func main() {

	//新建两个协程
	go person1()
	go person2()
	// 让主协程不退出
	for {
	}
}
