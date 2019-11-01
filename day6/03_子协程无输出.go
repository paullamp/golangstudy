package main

import (
	"fmt"
	"time"
)

func main() {
	// 子协程创建后，主协程立即退出。所以，无法输出任何信息
	go func() {
		fmt.Println("Hello world golang routine")
		time.Sleep(time.Second)
	}()
}
