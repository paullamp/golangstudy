package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "abc azc a7c aac 888 a9c"
	//1) 解释规则
	reg1 := regexp.MustCompile(`a.c`)
	if reg1 == nil {
		fmt.Println("regex error")
		return
	}
	//2) 根据规则提取关键字信息
	fmt.Println(reg1.FindAllStringSubmatch(buf, -1))
}
