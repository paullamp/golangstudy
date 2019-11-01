package main

import (
	"fmt"
	"strconv"
)

type mystu struct {
	Name string
	Age  int
}

func (m mystu) String() string {
	return "Name:" + m.Name + " Age:" + strconv.Itoa(m.Age)
}
func HandleChan(c chan mystu) {

	fmt.Println("HandleChan")
	c <- mystu{"zhangsan", 18}
}
func main() {
	c := make(chan mystu)
	go HandleChan(c)
	d := <-c
	fmt.Println("Function will finished in order")
	fmt.Println(d)
}
