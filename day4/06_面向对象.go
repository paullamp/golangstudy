package main

import "fmt"

func add01(a, b int) int {
	return a + b
}

type myint int

//num1 就是接收者，接收者传递了一个参数
//面向对象只是换了一种形式，叫成接收者。其实与python或java中的面向对象是一样的。
//如果使用结构体的形式，应该就能体会出来， 并且使用结构体的继承
func (num1 myint) add02(num2 myint) myint {
	return num1 + num2
}
func main() {
	a := myint(18)
	fmt.Println(a.add02(10))
}
