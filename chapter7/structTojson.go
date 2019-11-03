package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type ServerSlice struct {
	Servers []Server
}

func main() {
	ss := ServerSlice{}
	ss.Servers = append(ss.Servers, Server{"NginxWebServer", "10.129.8.21"})
	ss.Servers = append(ss.Servers, Server{"ShangHaiWebServer", "10.5.1.15"})
	jsondata, err := json.Marshal(ss)
	if err != nil {
		fmt.Println("output json data:", err)
		return
	}
	fmt.Println(string(jsondata))
	fmt.Println(strings.Repeat("---", 20))
	jsonpretty, err := json.MarshalIndent(ss, " ", "  ")
	if err != nil {
		fmt.Println("parse json pretty failed:", err)
		return
	}
	fmt.Println(string(jsonpretty))
}
