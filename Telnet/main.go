package main

import (
	"io"
	"os"
	"strings"

	// "strings"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

//获取http网页内容函数
func HttpGet(url string) (result string) {
	httpresp, errGet := http.Get(url)

	if errGet != nil {
		fmt.Println("http.Get读取网页错误:", errGet)
		return
	}

	buf := make([]byte, 1024*4)
	for {
		n, errBodyRead := httpresp.Body.Read(buf)
		if errBodyRead == io.EOF {
			break
		} else if errBodyRead != nil {
			fmt.Println("httpresp.Body.Read读取网页Body错误:", errBodyRead)
			return
		}
		result += string(buf[:n])
	}

	defer httpresp.Body.Close()
	//输出网页内容
	// fmt.Println(result)
	return result

}

func ScrapMainPage(n int, mainurl string, ch chan<- int) {
	fmt.Printf("开始第%d页抓取\n", n)
	onepage := map[string]string{}
	result := HttpGet(mainurl)

	/*
		从网页内容中读取需要匹配的子url
		需要抓取的内容格式
		<h1 class="dp-b"><a href="https://www.pengfue.com/content_1857656_1.html" target="_blank">
	*/
	subpagerestring := `<h1 class="dp-b"><a href="(?s:(.*?))" target="_blank">`

	subpagereg := regexp.MustCompile(subpagerestring)
	if subpagereg == nil {
		fmt.Println("正则表达式解析错误", subpagereg)
		return
	}

	subpagess := subpagereg.FindAllStringSubmatch(result, -1)
	//返回的结果是[][]string，　string的二维数组
	for _, v := range subpagess {
		// fmt.Println(v[1])
		title, content := ScrapSecondPage(v[1])
		onepage[title] = content
	}
	// fmt.Println(onepage)
	StoreToFile(n, onepage)
	ch <- n

}

func StoreToFile(n int, str map[string]string) {
	filename := "xiaohua_" + strconv.Itoa(n) + ".txt"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("os.Create error :", err)
		return
	}
	defer file.Close()
	for key, value := range str {
		file.WriteString(key + "\r\n")
		file.WriteString(value + "\r\n")
		file.WriteString("=================================================================================\r\n")
	}
}

//抓取子页面函数
func ScrapSecondPage(secondPageUrl string) (title, content string) {

	result := HttpGet(secondPageUrl)
	//只取第一个<h1></h1>里的mpw
	secondpageretitle := `<h1>(?s:(.*?))</h1>`
	secondpagerecontent := `<div class="content-txt pt10">(?s:(.*?))<a id="prev"`

	//读取标题
	secondPageTitleReg := regexp.MustCompile(secondpageretitle)
	if secondPageTitleReg == nil {
		fmt.Println("正则表达式解析错误", secondPageTitleReg)
		return
	}
	secondPageTitless := secondPageTitleReg.FindAllStringSubmatch(result, 1)
	// fmt.Printf("%s\n", strings.TrimSpace(secondPageTitless[0][1]))

	//读取内容
	secondPageContentReg := regexp.MustCompile(secondpagerecontent)
	if secondPageContentReg == nil {
		fmt.Println("正则表达式解析错误", secondPageContentReg)
		return
	}
	secondPageContentss := secondPageContentReg.FindAllStringSubmatch(result, -1)
	// fmt.Printf("%s\n", strings.TrimSpace(secondPageContentss[0][1]))
	title, content = strings.TrimSpace(secondPageTitless[0][1]), strings.TrimSpace(secondPageContentss[0][1])
	return
}

//抓取工具主要工作函数
func DoWork(startpage, endpage int) {
	ch := make(chan int)
	for i := startpage; i <= endpage; i++ {
		url := "https://www.pengfue.com/index_" + strconv.Itoa(i) + ".html"
		go ScrapMainPage(i, url, ch)
	}
	for i := startpage; i <= endpage; i++ {
		pageid := <-ch
		fmt.Printf("第 %d 页已存存储完毕\n", pageid)
	}
}

//golang 小爬虫，用于抓取网页的图片
func main() {
	//输入需要爬取的起始页
	var startpage, endpage int

	fmt.Printf("输入起始页:\t")
	fmt.Scan(&startpage)
	//输入需要爬取的终止页
	fmt.Printf("输入终止页:\t")
	fmt.Scan(&endpage)
	//输出信息
	fmt.Printf("需要抓取的页面范围是：%d 到 %d\n", startpage, endpage)
	//开始抓取

	DoWork(startpage, endpage)
}
