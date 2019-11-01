package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
		// ch <- 666 // channel 关闭后，不能再往里发送数据，不然会导致panic
	}()
	for num := range ch { // 当channel无数据后，会直接关闭channel。不需要如main01那样进行判断
		fmt.Println("num: ", num)
	}
}

func main01() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
		// ch <- 666 // channel 关闭后，不能再往里发送数据，不然会导致panic
	}()

	for {
		// ok 为true,说明通道没有关闭。如果为false，则说明通道关闭
		if num, ok := <-ch; ok {
			fmt.Printf("num = %d\n", num)
		} else {
			fmt.Println("ch is closed")
			break
		}
	}
}
