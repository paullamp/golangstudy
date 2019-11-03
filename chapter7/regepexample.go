package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	// 编写功能，获取一个URL网页内的所有的URL地址
	url := "http://www.baidu.com"
	resp, httpErr := http.Get(url)
	if httpErr != nil {
		fmt.Println("http error occured:", httpErr)
		return
	}

	content, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println("read resp.body failed:", readErr)
		return
	}
	patter := r"^((ht|f)tps?):\/\/[\w\-]+(\.[\w\-]+)+([\w\-.,@?^=%&:\/~+#]*[\w\-@?^=%&\/~+#])?$"
	reg, regerr := regexp.Compile(pattern)
	if regerr != nil {
		fmt.Println("compile reg failed:", regerr)
		return
	}
	ss := reg.FindAllString(string(content), -1)
	for _, s := range ss {
		fmt.Println(s)
	}
}
