package main //每个程序都必须要有一个主入口，同目录下的文件，只能归属于一个包，不同目录，可以有多个包

import (
	"fmt"
	"time"
)

type student struct {
	name   string
	age    int
	salary float64
}

func main() {
	// stu1 := map[string]int{
	// 	"name":   18,
	// 	"age":    20,
	// 	"salary": 22201,
	// }
	// fmt.Println(stu1)

	// stu2 := student{"zhangsan", 18, 198212.12}
	// fmt.Println(stu2)

	// var stu []*student // 刚定义，未分配空间无法存储数据
	// stu[0].name = "zhangsan"
	// stu := make([]*student, 0)
	// stu := make([]student, 3)
	// testgetSlicArgs(stu)
	// fmt.Println(stu)
	// fmt.Println("stu", stu)
	// s1 := student{"zhangsan", 18, 19.18}
	// testStruct(s1)
	// sint := []int{3, 5, 7, 12, 8}
	// for key, value := range sint {
	// 	fmt.Println(key, value)
	// }
	// s1 := map[string]int{
	// 	"zhagnsan": 10,
	// 	"lisi":     20,
	// 	"wangwu":   30,
	// }
	// testmap(s1)
	// s1 := make(map[string]int)
	// if s1 == nil {
	// 	fmt.Println("nil")
	// 	fmt.Println(s1)
	// }

	// s := []int{}
	// fmt.Println(len(s), cap(s))

	a1 := time.Now()
	time.Sleep(5)
	a2 := time.Now()
	if a1.Before(a2) {
		fmt.Println(a2)
	}
}

func testgetSlicArgs(stu []student) {
	// stu = make([]student, 3)
	// stu = append(stu, student{name: "zhangsan"})
	// stu = append(stu, student{name: "work"})
	stu[0].name = "zhangsan"
	stu[1].name = "lisi"
	fmt.Println("stu in testgetsliceargs ", stu)
	// for key, value := range stu {
	// 	fmt.Println(key, value)
	// }

}

func testStruct(s1 student) {
	fmt.Println(s1)
}

func testmap(s1 map[string]int) {
	for k, v := range s1 {
		fmt.Println(k, v)
	}
}
