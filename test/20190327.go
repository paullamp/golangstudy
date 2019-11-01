package main

import (
	"fmt"
)

func runloop() {
	for i := 0; i < 10; i++ {
		fmt.Println("number is :", i)
		// ch <- 10
	}
}

func main() {
	// str := make(chan string)
	// // go func(input string) {
	// // 	str <- input
	// // }("Hello world")
	// str <- "Hello world" // 无缓冲信道的读或写，都会阻塞当前的协程。若无程序取走当前协程里的内容，它将不再往下执行。
	// fmt.Println("channel message : ", <-str)

	str := make(chan string, 2)
	str <- "Hello world"
	fmt.Println("channel message have two slot:", <-str)
}
