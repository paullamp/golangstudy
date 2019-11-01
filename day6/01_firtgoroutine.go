package main

import (
	"fmt"
	"time"
)

func newTask() {
	for {
		fmt.Println("This is new task")
		time.Sleep(time.Second)
	}
}
func main() {
	go newTask()
	for {
		fmt.Println("this is main goroutine")
		time.Sleep(time.Second)
	}
}
