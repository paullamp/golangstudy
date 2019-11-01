package main

import (
	"fmt"
)

func test() {
	fmt.Println("welcome to golang")
	return
	fmt.Println("footer in test")
}

func call2() {
	fmt.Println("call2 function:before call test")
	test()
	fmt.Println("call2 function:after call test")
}
func main() {
	fmt.Println("befere call2 function call")
	call2()
	fmt.Println("after call2 function call")
}
