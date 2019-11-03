package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("workhard", "or"))
	fmt.Println(strings.Split("china-beijing-shangehai-xian", "-"))
	fmt.Println(strings.Repeat("gogo", 3))
	fmt.Println(strings.TrimSpace("  work   "))

	fmt.Println(strconv.Itoa(1987))
	datavalue, converr := strconv.Atoi("20181105")
	if converr != nil {
		fmt.Println("strconv to int failed", converr)
		return
	}

	fmt.Println(datavalue)
}
