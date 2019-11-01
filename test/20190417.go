package main

import (
	"fmt"
)

func comsume(a []chan int) {
	for ch := range a {
		fmt.Println(ch)
	}

}

func main() {
	c := make([]chan int, 10)
	comsume(c)
}
