package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Just run test")
}

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("div 0")
	}
	return a / b, nil
}
