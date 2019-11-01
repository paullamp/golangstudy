package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	Person
	dept string
}

func (this Person) PrintInfo() {
	fmt.Println("this is person 's printinfo func")
}

func (this Student) PrintInfo() {
	fmt.Println("this is Student' printinfo func")
}

func main() {
	s := Student{Person{"zhagnsna", 19}, "computer since"}
	//person中有printinfo方法， 现在是使用student的方法，重写后的。
	s.PrintInfo()
	//显式调用person方法
	s.Person.PrintInfo()
}
