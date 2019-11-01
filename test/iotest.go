package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("server.go")
	if err != nil {
		log.Fatal("read content failed")
	}
	fmt.Println(string(content))
}
