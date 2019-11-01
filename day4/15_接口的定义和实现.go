package main

import "fmt"

type Humber interface {
	SayHi()
}
type Person struct {
	name string
	age  int
}

func (p Person) SayHi() {
	fmt.Printf("Person:%s,%d say hi\n", p.name, p.age)
}

type mystr string

func (st1 mystr) SayHi() {
	fmt.Printf("mystr :%s\n", st1)
}

// 多态，同一个方法，多种不同的输出。 依赖于传入接口的具体类型
func whosayhi(i Humber) {
	i.SayHi()
}
func main() {
	p1 := Person{"zhangsna", 22}
	var s1 mystr = "lisi"
	whosayhi(p1)
	whosayhi(s1)

	s := make([]Humber, 2)
	s[0] = p1
	s[1] = s1
	for _, value := range s {
		value.SayHi()

	}
}
func main01() {
	var i Humber
	i = Person{"zhangsan", 19}
	i.SayHi()
	var s1 mystr = "hello world"
	i = s1
	i.SayHi()
}
