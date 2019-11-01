package main

import "fmt"
import "strconv"

func main() {
	fmt.Println(strconv.Atoi("8899"))
	fmt.Println(strconv.Itoa(8893))
	fmt.Println(strconv.ParseBool("true"))
	var b bool = true
	fmt.Println(strconv.FormatBool(b))
}
