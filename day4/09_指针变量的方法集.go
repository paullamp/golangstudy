package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (this Person) SetInfoValue() {
	fmt.Println("person values function")
}

func (this *Person) SetInfoPointer() {
	fmt.Println("person Pointer function")
}
func main() {
	var p Person = Person{"mkie", 18}
	//p是一个值，非指针，可以直接调用以下两个函数
	p.SetInfoPointer()
	p.SetInfoValue()

	//(&p)是一年指针，非常量，同样可以调用以下两个函数
	(&p).SetInfoPointer()
	(&p).SetInfoValue()
}
