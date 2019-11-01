package main

import (
	"fmt"
)

//此通道只能写，不能读
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

//此通道只能读，不能写
func comsumer(in <-chan int) {
	for num := range in {
		fmt.Printf("num is :%d \n", num)
	}
}
func main() {
	//创建一个双向chan用于接收，发送数据
	ch := make(chan int) //注意，需要使用make创建，不能直接使用var c chan int

	//创建一个生产者进程
	go producer(ch)

	//创建一个消费者函数，直接在主进程中运行
	comsumer(ch)
}
