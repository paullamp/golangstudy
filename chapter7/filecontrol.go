package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fd, err := os.Open("j2.go")
	if err != nil {
		fmt.Println("open j2.go failed", err)
		return
	}
	rd := bufio.NewReader(fd)
	for {
		s, err := rd.ReadString('\n')
		fmt.Print(s)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read string failed")
			return
		}
	}
	defer fd.Close()

}
