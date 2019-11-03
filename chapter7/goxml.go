package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

type ServerSlice struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Servers     []Server `xml:"server"`
	Description string   `xml:",innerxml"`
}

func main() {
	fd, openerr := os.Open("example.xml")
	if openerr != nil {
		fmt.Println("Open xml failed", openerr)
		return
	}
	fileread, err := ioutil.ReadAll(fd)
	if err != nil {
		fmt.Println("read fd failed", err)
		return
	}
	v := new(ServerSlice)
	err = xml.Unmarshal(fileread, v)
	if err != nil {
		fmt.Println("unmarshal file failed:", err)
		return
	}
	// fmt.Println(v.Servers)
	for _, value := range v.Servers {
		fmt.Println(value.ServerName, value.ServerIP)
	}
}
