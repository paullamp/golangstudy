package main

import "fmt"
import "strings"

func main() {
	str := "myname@is@zhangsan@lisi"
	s := strings.Split(str, "@")
	for _, data := range s {
		fmt.Println(data)
	}
	str2 := "    Hello     wrold     "
	fmt.Println(strings.Trim(str2, " "))
}
