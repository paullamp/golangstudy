package main

import (
	"fmt"

	"github.com/astaxie/goredis"
)

func main() {
	client := goredis.Client{Addr: "127.0.0.1:8888"}

	client.Set("name", []byte("welldone"))
	val, err := client.Get("name")
	if err != nil {
		fmt.Println("Get value failed:", err)
		return
	}
	fmt.Println(string(val))
}
