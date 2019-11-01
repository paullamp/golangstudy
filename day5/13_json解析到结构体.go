package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string   `json:"company"`
	Subject []string `json:"subject"` //重定义需要解析的名称
	Isok    bool     `json:"isok"`
	Price   int      `json:"price"`
}

//如果只需要其中的一个字段，那么只需要定义有一个字段的结构体
type IT2 struct {
	Isok bool `json:"isok"`
}

func main() {
	jsonbuf := `{"company":"Bjglodon","subject":["go","c++","erlang","php","bash"],"isok":true,"price":99}`
	var it IT
	//解析到结构体
	err := json.Unmarshal([]byte(jsonbuf), &it)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(it)
	fmt.Printf("%+v\n", it)

	var it2 IT2
	err = json.Unmarshal([]byte(jsonbuf), &it2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", it2)
}
