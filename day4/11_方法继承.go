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

func main() {
	s := Student{Person{"zhagnsna", 19}, "computer since"}
	//student 类型无printinfo方法。此方法为从person继承
	s.PrintInfo()
}
