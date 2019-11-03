package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	Name     string
	Serverip string
}

type serverslice struct {
	Servers []Server
}

func main() {
	// s := serverslice{}
	var s serverslice
	s.Servers = append(s.Servers, Server{"bj-dhcp-server", "192.168.0.201"})
	s.Servers = append(s.Servers, Server{"sh-dhcp-server", "10.0.5.250"})

	jsonbyte, err := json.Marshal(s)
	if err != nil {
		fmt.Println("parse to json failed")
		// return
	}
	fmt.Println(string(jsonbyte))
}
