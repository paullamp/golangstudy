package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)
	timer.Reset(1 * time.Second)
	<-timer.C
	fmt.Println("time is up")
}
func main01() {
	timer := time.NewTimer(3 * time.Second)
	timer.Stop()
	<-timer.C
	fmt.Println("time is up")
}
