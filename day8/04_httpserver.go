package main

import (
	"fmt"
	"net/http"
)

func Handlemike(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.Body = ", r.Body)
	fmt.Println("r.Method = ", r.Method)
	fmt.Println("r.RemoteAddr =", r.RemoteAddr)
	fmt.Println("r.header = ", r.Header)
	fmt.Println("r.URL = ", r.URL)
	w.Write([]byte("Hello mike"))
}

func main() {
	//注册url处理函数
	http.HandleFunc("/mike.html", Handlemike)

	//启动监听及服务
	http.ListenAndServe("127.0.0.1:8000", nil)
}
