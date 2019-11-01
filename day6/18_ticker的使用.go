package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second) // 定时执行
	i := 0
	for {
		m := <-ticker.C
		fmt.Println(m)
		i++
		fmt.Println("i=", i)
		if i == 5 {
			ticker.Stop()
			break
		}
	}
}
