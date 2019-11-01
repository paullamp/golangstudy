package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	fmt.Printf("len(ch)=%d cap(ch)=%d\n", len(ch), cap(ch))

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("newTask[%d], len(ch)=%d , cap(ch)=%d\n", i, len(ch), cap(ch))
			ch <- i
		}
	}()

	time.Sleep(time.Second * 2)
	for i := 0; i < 5; i++ {
		num := <-ch
		fmt.Printf("Main task:%d\n", num)
	}
}
