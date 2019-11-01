package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("cccccccccccccccccccccc")
	// return           //终止此函数
	runtime.Goexit() //终止所在的协程
	fmt.Println("dddddddddddddddddddd")
}
func main() {
	//一个goroutine
	go func() {
		fmt.Println("aaaaaaaaaaaaaaaaaaaaaaa")
		test()
		fmt.Println("bbbbbbbbbbbbbb")
	}()
	//写一个死循环，保证主协程不退出
	for {
	}
}
