package main

import (
	"fmt"
	"regexp"
)

func main() {
	number := "13566891244"
	statu, compareErr := regexp.MatchString("^[0-9]+$", number)
	if compareErr != nil {
		fmt.Println("compare error:", compareErr)
		return
	}
	if statu {
		fmt.Println("Matched")
	} else {
		fmt.Println("Un-Matched")
	}
}
