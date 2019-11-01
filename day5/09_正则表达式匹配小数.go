package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "1.34 abc 88.12 7. efsf 112.42"
	//定义正则表达式, + 匹配前一个字符的一个或多个
	reg1 := regexp.MustCompile(`\d+\.\d+`)
	//查找模式
	res := reg1.FindAllStringSubmatch(buf, -1)

	for key, value := range res {
		fmt.Println(key, value)
	}
}
