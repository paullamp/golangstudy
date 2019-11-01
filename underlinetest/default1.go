package underlinetest

import (
	"fmt"
)

var num1 int = 5
var num2 int = 10

func main() {
	fmt.Println("Hello main function in underlinetest")

}

func init() {
	fmt.Println("This is init function in underlinetest")
	fmt.Println("Two number in package 全局变量:", num1, num2)
}

func AddTwo(a, b int) int {
	return a + b
}
