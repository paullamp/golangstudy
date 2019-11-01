package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  byte
}

//类似php或python中的this 写法
func (this Person) printInfo() {
	fmt.Printf("%+v\n", this)
}

func (this *Person) SetInfo(n string, a int, s byte) {
	this.name, this.age, this.sex = n, a, s
}

type nint int

// add a method to int type , 原生类型是不可以的，必须自定义一次。
func (this nint) add(other nint) nint {
	return this + other
}

type long int

func (this long) test() {}

type char byte

func (this char) test() {}

func main() {
	p := Person{"zhangsan", 18, 'M'}
	p.printInfo()
	p1 := new(Person)
	p1.SetInfo("zhangsan", 19, 'M')
	fmt.Println(p1)
}
