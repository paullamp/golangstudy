package main

import (
	"fmt"
)

func main() {
	// fmt.Println("back home from 20190325")
	// var ch chan int
	ch := make(chan int)
	ch <- 5
	// go func() {
	// 	fmt.Println("add a number to ch")
	// 	ch <- 10
	// }()

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("welcome to %d\n", i)
	// }
	d, ok := <-ch
	fmt.Println(ok)
	fmt.Printf("chan ch is : %d\n", d)
}
