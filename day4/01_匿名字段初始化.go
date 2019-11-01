package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person //只有类型，没有名字，匿名字段,继承了person里面的成员
	id     int
	addr   string
}

func main() {
	s1 := new(Student)
	s1.age = 18
	s1.name = "zhangsna"
	s1.sex = 'M'
	fmt.Println(s1)
	// 顺序初始化
	var s2 Student = Student{Person{"mkie", 'M', 18}, 1888812, "Hello street world"}

	fmt.Println(s2)
	//自动推导类型
	s3 := Student{Person{"susan", 'S', 20}, 1988232, "Beijing haidian"}
	fmt.Println(s3)
	//%+v显示更详细信息
	fmt.Printf("%+v\n", s3)
	//指定成员初始化，没有初始化的成员自动赋值0
	s4 := Student{id: 1}
	fmt.Printf("%+v\n", s4)
}
