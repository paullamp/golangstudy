package main

import (
	"encoding/json"
	"fmt"
)

type mytest struct {
	Name string
	Age  string
	Dept string
}

func main() {
	a := make(map[string]string)
	a["name"] = "zhangsan"
	a["age"] = "12"
	a["dept"] = "it"
	fmt.Println(a)
	myjson, err := json.Marshal(a)
	if err != nil {
		fmt.Println("解析成json出错:", err)
		return
	}
	// 将json转换成map,通过make(map[string]string)来创建一个对象，可以直接存值。
	// 是一个值的变量，可直接存值
	// b := make(map[string]string)
	// 或是直接使用{} ，　定义一个空的值。可以直接存
	// b := map[string]string{}

	//使用结构体接收json的数据
	b := mytest{}
	if err := json.Unmarshal(myjson, &b); err != nil {
		fmt.Println("将json的[]byte切片转换成字典的错误:", err)
		return
	}
	fmt.Println(b)
	fmt.Println("Age:", b.Age)
	fmt.Println("Name:", b.Name)
	fmt.Println("Dept:", b.Dept)

	c := []mytest{}
	//若是切片的len为0,cap为0,肯定是不能存储数据
	fmt.Println("Len is :", len(c), "Cap is:", cap(c))
	// c[0] = mytest{"zhangsna", "25", "yanfabu"}  -- 报错，提示index out of range
	c = append(c, mytest{"hanwangwang", "30", "caiwubu"})
	fmt.Println(c)
}
