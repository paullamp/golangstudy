package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) PrintInfo() {
	fmt.Println("this is person 's printinfo func")
	fmt.Printf("PrintInfo:address is : %p,value is : %v\n", &p, p)
}
func (p *Person) SetInfo() {
	fmt.Println("this is person 's setinfo func")
	fmt.Printf("SetInfo:address is : %p,value is : %v\n", p, *p)
}
func main() {
	p := Person{"zhagnsan", 109}
	fmt.Printf("main: address is : %p,value is : %v\n", &p, p)
	// p.PrintInfo()
	vfunc := p.PrintInfo  //方法值
	pfunc := (&p).SetInfo // 方法值
	vfunc()
	pfunc()
}
