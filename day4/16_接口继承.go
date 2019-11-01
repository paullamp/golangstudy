package main

import "fmt"

type Humber interface {
	SayHi()
}
type Personer interface {
	Humber
	sing(lrc string)
}

type Student struct {
	name string
	age  int
}

func (stu *Student) SayHi() {
	fmt.Printf("Stu: %s, %d\n", stu.name, stu.age)
}

func (stu *Student) sing(lrc string) {
	fmt.Printf("Stu:%s, %d, is singing : %s\n", stu.name, stu.age, lrc)
}

func main() {
	var i Personer
	i = &Student{"mkie", 18}
	// i.SayHi()
	i.sing("welldone")
}
