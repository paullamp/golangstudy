package main

import (
	// "fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/go", goindex)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

//func(ResponseWriter, *Request)
func goindex(respose http.ResponseWriter, req *http.Request) {
	respose.Write([]byte("Helloworld"))
}
