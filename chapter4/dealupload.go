package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t, err := template.ParseFiles("upload.html")
		if err != nil {
			fmt.Println("parsefile failed")
			return
		}
		t.Execute(w, nil)
	}

	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		file, handle, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println("read form file failed")
			return
		}
		defer file.Close()
		f, err := os.Create(handle.Filename)
		if err != nil {
			fmt.Println("create local file failed")
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
func main() {
	http.HandleFunc("/upload", handleLogin)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println("build http service failed")
		return
	}
}
