package main

import (
	"fmt"
	"net/http"

	"github.com/drone/routes"
)

func getuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are get user:%s", uid)
}

func modifyuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are modify user: %s", uid)
}

func main() {
	mux := routes.New()
	mux.Get("/user/:uid", getuser)
	mux.Post("/usr/:uid", modifyuser())
	http.Handle("/", mux)
	http.ListenAndServe(":9999", nil)
}
