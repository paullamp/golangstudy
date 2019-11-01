package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  byte
}

type Student struct {
	*Person
	int
	string
}

func main() {
	s1 := Student{&Person{"zhangsna", 18, 'M'}, 22323, "SZ"}
	fmt.Println(s1)
	fmt.Println(s1.name)

	//直接使用NEW分配指针空间
	var s2 Student
	s2.Person = new(Person)
	s1.age = 19
	s2.sex = 'F'
	s2.name = "GOGO"
	fmt.Println(s2.Person)
}
