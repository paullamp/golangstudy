package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string   `json:"-"`       //冒号前后不能有空格，此写法表明，此字段不解析成json
	Subject []string `json:"subj"`    // 将字段重命名
	Isok    bool     `json:",string"` // 将其他类型的值转换成字符串类型
	Price   int
}

func main() {
	m := make(map[string]interface{}, 4)
	m["Company"] = "itcast"
	m["Subject"] = []string{"golang", "c++", "vb", "php"}
	m["Isok"] = true
	m["price"] = 88 //map格式下，小写同样可以导出

	// result, err := json.Marshal(m)
	result, err := json.MarshalIndent(m, "", "	") //格式化输出
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(result))
}
