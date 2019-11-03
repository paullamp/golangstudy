package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name   string `json:"Name"`
	Age    int    `json:"Age"`
	Salary int    `json:"Salary"`
}

type Students struct {
	Stus []Student
}

func main() {
	// s := Students{}
	// s.Stus = []Student{Student{"zhangsan", 23, 1800}, Student{"lisi", 23, 1800}}
	// jsonbytes, err := json.Marshal(s)
	// if err != nil {
	// 	fmt.Println("json Marshal failed:", err)
	// 	return
	// }

	s := Student{"zhangsan", 23, 1800}
	jsonbytes, err := json.Marshal(s)
	if err != nil {
		fmt.Println("parse struct to json failed")
		return
	}
	fmt.Println("In struct type: ", s)
	fmt.Println("In Json type:", string(jsonbytes))
}
