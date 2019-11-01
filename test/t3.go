package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := map[string]interface{}{}
	a["name"] = "zhangsan"
	a["age"] = 18
	a["salary"] = 3306
	fmt.Println(reflect.TypeOf(a["name"]))
	fmt.Println("通过断言方式进行处理:", a["name"].(string))
	// fmt.Println("通过断言方式进行处理:", a["name"].(int))
	//　通常情况下，断言处理需要使用value,ok进行接收，　通过ok来进行类型的判断，不然，如果断言错误
	//　会导致程序panic
	if value, ok := a["name"].(int); ok {
		fmt.Println("断言为int 类型成功，值为：", value)
	} else {
		fmt.Println("断言成int类型失败")
	}

	//使用switch case 处理断言
	var testalltype interface{}
	testalltype = "helloworld"
	switch testalltype.(type) {
	case string:
		fmt.Println("String match")
		fmt.Println("测试字符串连接操作：", testalltype.(string)+" BeijingNew")
	case int:
		fmt.Println("Int match")
	default:
		fmt.Println("Nothing match")
	}
	fmt.Println("nil 类型测试")
	var num float64 = 1.24324
	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)
	numType := reflect.TypeOf(num)
	fmt.Println("Pointer is :", pointer)
	fmt.Println("Value is:", value)
	fmt.Println("num type is : ", numType)
}
