package main

import "fmt"

func main() {
	type Person struct {
		name string
		age  int
		sex  byte
	}

	type Student struct {
		Person
		int
		string
	}

	s1 := Student{Person{"zhangsna", 18, 'M'}, 10, "Hello street"}
	s2 := Student{Person{"zhangsan", 18, 'F'}, 19, "sz"}
	fmt.Printf("%+v\n", s1)
	fmt.Printf("%+v\n", s2)
	var s3 Student
	s3.Person = Person{"zhangsna", 19, 'F'}
	s3.int = 18
	s3.string = "Helhehe"
	fmt.Println(s3)
}
