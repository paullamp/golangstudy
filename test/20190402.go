package main

import (
	"fmt"
)

func TestByte(dst []byte) {
	src := []byte("hello china")
	dst = src
	fmt.Println("len(src):", len(src), "cap(src)", cap(src), src)
	fmt.Println("len(dst):", len(dst), "cap(dst)", cap(dst), dst)
	return
}
func main() {
	var dst []byte

}
