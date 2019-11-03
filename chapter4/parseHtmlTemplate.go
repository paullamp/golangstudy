package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	t, teer := template.ParseFiles("index.html")
	if teer != nil {
		fmt.Println("parse index failed")
		return
	}
	t.Execute(w, nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	t, parseTemplateErr := template.ParseFiles("form.html")
	if parseTemplateErr != nil {
		fmt.Println("Parse template failed")
		return
	}
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/login", handleLogin)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println("listen and serve in http server failed")
		return
	}
}
