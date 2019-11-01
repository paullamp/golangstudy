package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 0)
	fmt.Printf("len(ch) = %d , cap(ch) = %d\n", len(ch), cap(ch))
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("new task: %d\n", i)
			ch <- i
		}
	}()
	for i := 0; i < 3; i++ {

		num := <-ch
		fmt.Printf("Main task:%d\n", num)
	}
}
