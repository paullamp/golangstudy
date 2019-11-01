package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  byte
}

//值传递，传递的是值的拷贝副本
func (p Person) SetInfoValue(n string, a int, s byte) {
	p.name = n
	p.age = a
	p.sex = s
	fmt.Printf("SetInfoValue Method : --- address is : %p, value is : %v\n", &p, p)
}

//引用传递，传递的是地址。
func (p *Person) SetInfoPointer(n string, a int, s byte) {
	p.name = n
	p.age = a
	p.sex = s
	fmt.Printf("SetInfoPointer Method : --- address is : %p, value is : %v\n", p, *p)
}

func main() {
	var p1 Person = Person{"Mike", 'F', 18}
	fmt.Printf("address is : %p, value is : %v\n", &p1, p1)
	p1.SetInfoValue("zhangsan", 'M', 25)
	(&p1).SetInfoPointer("lisi", 'M', 88)
	fmt.Printf("address is : %p, value is : %v\n", &p1, p1)
}
