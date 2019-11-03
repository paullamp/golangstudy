package main

import (
	"fmt"
	"io"

	// "net"
	"net/http"
)

func main() {
	resp, geterr := http.Get("http://www.baidu.com")
	if geterr != nil {
		fmt.Println("geterror:", geterr)
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 1024)
	for {
		n, readerr := resp.Body.Read(buf)
		fmt.Println(string(buf[:n]))
		if readerr == io.EOF {
			break
		}
		if readerr != nil {
			fmt.Println("readerror: ", readerr)
			return
		}
	}

}
