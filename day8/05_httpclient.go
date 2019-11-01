package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http.get error = ", err)
		return
	}
	defer resp.Body.Close()

	// 打印返回的状态码
	fmt.Println(resp.Status)
	bodycontent := make([]byte, 1024*4)
	var tmp string
	for {
		n, resp_body_read_err := resp.Body.Read(bodycontent)
		if n == 0 {
			// fmt.Println("read response body error = ", resp_body_read_err)　//不需要输出此行
			break //如果换成return ,会提前中止函数，无法进入后面的打印进程
		}
		tmp += string(bodycontent[:n])
	}

	fmt.Println("tmp = ", tmp)
}
