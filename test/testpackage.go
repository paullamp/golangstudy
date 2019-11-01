package main

import (
	// "fmt"
	// "underlinetest" // 只要在import的包列表中， 那么包中所包含的 func init()都会执行
	_ "underlinetest"
)

func main() {
	// fmt.Println(underlinetest.AddTwo(3, 5)) // 此时，underlinetest在主函数中不可调用
	// fmt.Println("Just undlerlinetest.init() will be run")
}
