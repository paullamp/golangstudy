package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonbuf := `{"company":"Bjglodon","subject":["go","c++","erlang","php","bash"],"isok":true,"price":99.88}`
	m := make(map[string]interface{}, 4)
	err := json.Unmarshal([]byte(jsonbuf), &m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m)
	//通过类型断言来获取内容
	for _, data := range m {
		switch value := data.(type) {
		case string:
			var str string
			str = value
			fmt.Printf("str is : %s\n", str)
			fmt.Printf("type string %s\n", value)
		case float64:
			fmt.Printf("type int : %v\n", value)
		case bool:
			fmt.Printf("type bool: %v\n", value)
		case []string:
			fmt.Printf("string type %v", value)
		case []interface{}:
			fmt.Printf("interface{} type %v", value)
		default:
			fmt.Println("nothing")
		}
	}

}
