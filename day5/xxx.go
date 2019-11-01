package main

import "fmt"

type stu struct {
	name string
	age  int
}

func main() {
	// var s1 stu
	s1 := new(stu)
	s1.name = "zhgnsna"
	s1.age = 18
	fmt.Println(s1)
}
