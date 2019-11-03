package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fd, err := os.Open("form.html")
	if err != nil {
		fmt.Println("open form failed")
		return
	}
	mess, err := ioutil.ReadAll(fd)
	if err != nil {
		fmt.Println("read content error")
		return
	}

	w.Write(mess)
}
func handleLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Write([]byte(r.PostFormValue("userpassword")))
}
func main() {
	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/login", handleLogin)
	httpListenErr := http.ListenAndServe("0.0.0.0:9090", nil)
	if httpListenErr != nil {
		fmt.Println("bind httpd service failed")
		return
	}
}
