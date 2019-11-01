package main

import (
	"fmt"
	"time"
)

func main() {
	// 主协程先退出后，其他子协程也跟着退出

	go func() {
		i := 0
		for {
			fmt.Printf("new gorouting:%d\n", i)
			i++
			time.Sleep(time.Second)
		}

	}()
	i := 0
	for {
		fmt.Printf("main goroutine: %d\n", i)
		i++
		time.Sleep(time.Second)
		if i == 2 {
			break
		}

	}
}
