package main

import (
	"encoding/json"
	"fmt"
)

/*
{
"company":"itcast",
"subject":[
"go",
"c++",
"python",
"Test"
],
"isok":true,
"price":6666
}
*/
// 如果想转成json格式， 所有结构体成员的首字母必须要大写，不然无法转换
// type IT struct {
// 	Company string
// 	Subject []string
// 	Isok    bool
// 	Price   int
// }

// tag 使用，二次编码
type IT struct {
	Company string   `json:"-"`       //冒号前后不能有空格，此写法表明，此字段不解析成json
	Subject []string `json:"subj"`    // 将字段重命名
	Isok    bool     `json:",string"` // 将其他类型的值转换成字符串类型
	Price   int
}

func main() {
	it := IT{"Bjglodon", []string{"go", "c++", "erlang", "php", "bash"}, true, 99}
	//{"Company":"Bjglodon","Subject":["go","c++","erlang","php","bash"],"Isok":true,"Price":99}
	// jsonres, err := json.Marshal(it)
	jsonres, err := json.MarshalIndent(it, " ", "    ") // 格式化显示
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonres))
}
