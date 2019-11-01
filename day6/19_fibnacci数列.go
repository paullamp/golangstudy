package main

import (
	"fmt"
)

func fab(ch chan<- int, flag <-chan bool) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case ok := <-flag:
			fmt.Println(ok)
			return
		}
	}
}
func main01() {
	//创建两个channel,一个用于接收数字，一个用于接收flag
	ch := make(chan int)
	flag := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			//必须将ch通道中的值取出，不然无法正常子协程会卡住
			number := <-ch
			fmt.Println(number)
		}
		flag <- true
	}()

	fab(ch, flag)
}

func main() {
	//创建两个channel,一个用于接收数字，一个用于接收flag
	ch := make(chan int)
	flag := make(chan bool)
	//为什么这样不行
	go fab(ch, flag) //需要放到前面，如果后面for循环后面，那么前面的会一直阻塞，造成死锁
	for i := 0; i < 10; i++ {
		//必须将ch通道中的值取出，不然无法正常子协程会卡住
		number := <-ch
		fmt.Println(number)
	}
	flag <- true

}
