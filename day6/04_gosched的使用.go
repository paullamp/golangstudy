package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go")
		}
	}()

	for i := 0; i < 2; i++ {
		//主协程出让时间片,让子协程先执行完毕
		runtime.Gosched()
		fmt.Println("Helloworld")
	}
}
