package main

import (
	"fmt"
	"io"
	"os"

	// "net"
	"net/http"
	"strconv"
)

//抓取内容的函数，结果保存到字符中中，　如果有错误，返回错误，无错误，返回nil
func HttpGet(url string) (result string, err error) {
	//通过　http.Get方法获取url内容
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		fmt.Println("Http.Get err = ", err1)
		return
	}
	defer resp.Body.Close()

	//读取resp.Body的内容
	buf := make([]byte, 1024*4)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 { //有可能读取出错，或是读取到结束
			if err == io.EOF {
				fmt.Println("文件读取结束")
			} else {
				fmt.Println("resp.body.read err = ", err)
			}
			break
		}
		result += string(buf[:n])
	}
	return
}

func DoWork(start, end int) {
	fmt.Printf("需要抓取的内容是：从%d到%d\n", start, end)
	/*
		https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=50
		https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=100
		https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=150
	*/
	for i := start; i <= end; i++ {
		//拼装url
		url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" +
			strconv.Itoa((i-1)*50)
		fmt.Println(url)

		//针对url进行抓取
		result, err := HttpGet(url)
		if err != nil {
			fmt.Println("HttpGet error = ", err)
			continue
		}
		//将抓取的内容写入文件
		fileName := strconv.Itoa(i) + ".html"
		f, err1 := os.Create(fileName)
		if err1 != nil {
			fmt.Println("os.Create error = ", err1)
			continue
		}

		f.Write([]byte(result))
		f.Close()

	}
}
func main() {
	//定义需要抓取的起始页和结束页变量，并且从终端进行输入
	var start, end int
	fmt.Println("请输入起始页：")
	fmt.Scan(&start)
	fmt.Println("请输入结束页：")
	fmt.Scan(&end)

	//开始抓取内容
	DoWork(start, end)
}
