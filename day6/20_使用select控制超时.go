package main

import (
	"fmt"
	"time"
)

func main() {
	//定义一个channel，接收数据，但是无数据传入
	ch := make(chan int)
	//定义一个关闭标志
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-ch:
				// 一直没有数据输入，所以轮循到后面的超时流程
				fmt.Println("a number in")
			case <-time.After(3 * time.Second):
				fmt.Println("timeout")
				quit <- true
			}
			//测试一下是否会执行到这一段代码，如果执行到，则表明执行过多次select验证
			//验证表明，只有当一次case生效，才会执行到这一段，不然会一直阻塞
			fmt.Println("select run once")
		}
	}()
	<-quit
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	fmt.Println("finished")
}
