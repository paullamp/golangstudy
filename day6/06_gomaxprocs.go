package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(1) //返回之前是工作在几个核心模式下,多核模式下，任务可以交替运行
	fmt.Println("the old process is :", n)
	for {
		// go func() {
		// 	fmt.Println(1)
		// }()
		go fmt.Println(1)
		fmt.Println(2)
	}
}
