package main

import (
	"fmt"
)

var a map[string]int = make(map[string]int, 10)

func main() {
	a["name"] = 10
	fmt.Println(a)
	b := "hello"
	fmt.Println(len(b))
}
