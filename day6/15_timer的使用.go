package main

import (
	"fmt"
	"time"
)

//验证时间到了，是不是还继续执行，还是只执行一次
//通过timer实现延时功能
func main() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("current time is :", time.Now())
	// t := <-timer.C
	// fmt.Println("timer time is :",t)
	fmt.Printf("len(timer.C) is %d, cap(timer.C) is %d\n", len(timer.C), cap(timer.C))
	for t := range timer.C {
		fmt.Println("timer time is :", t)
	}
}

func main01() {
	timer := time.NewTimer(time.Second * 2)
	fmt.Println("Current time is : ", time.Now())
	//2s后，会往timer.C中写数据，没有数据前会阻塞
	t := <-timer.C //获取定时器管道内的内容
	fmt.Println("timer time is :", t)
}
