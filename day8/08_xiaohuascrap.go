package main

/*
1.需要抓取的url规律：
https://www.pengfue.com/xiaohua_1.html
https://www.pengfue.com/xiaohua_2.html  页面递增

2.需抓取的主页规律
<h1 class="dp-b"><a href=" URL "  抓取具体内容的URL

3.单笑话页
<h1>标题</h1>　　只取第一个
<div class="content-txt pt10">　内容 <a id="prev"

*/

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//抓取页面
func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		fmt.Println("http.get error = ", err1)
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 1024*4)
	//定义一个死循环不停读取内容,存取到result中
	for {
		n, err2 := resp.Body.Read(buf)
		if err2 != nil {
			fmt.Println("resp.Body.Read error = ", err2)
			break
		}
		result += string(buf[:n])
	}

	return
}

//保存每一页的内容到文件
func StoreToFile(i int, fileTitles, fileContents []string) {
	f, err := os.Create("xiaohua_" + strconv.Itoa(i) + ".txt")
	if err != nil {
		fmt.Println("os.create err = ", err)
		return
	}
	defer f.Close()
	for index := 0; index < len(fileTitles); index++ {
		f.WriteString(fileTitles[index] + "\r\n")
		f.WriteString(fileContents[index] + "\r\n")
		f.WriteString("=================================\r\n")
	}
}

//爬取主页面
func SpiderMainPage(i int, pagechan chan<- int) {
	url := "https://www.pengfue.com/xiaohua_" + strconv.Itoa(i) + ".html"
	fmt.Println("url = ", url)

	//开始抓取页面
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("httpGet err = ", err)
		return
	}
	// fmt.Println("r=", result)
	//从抓取的结果中获取关键信息
	re1 := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	if re1 == nil {
		fmt.Println("regexp.MustCompile err")
		return
	}
	ss := re1.FindAllStringSubmatch(result, -1)
	fileTitles := make([]string, 0)   //定义字符切片，用于存储所有的标题
	fileContents := make([]string, 0) //定义字符切片，用于存储所有的内容
	for _, val := range ss {
		// fmt.Println(val[1])
		onePageUrl := val[1]
		title, content, err2 := SpiderOnePage(onePageUrl)
		if err2 != nil {
			fmt.Println("SpiderOnepage err=", err2)
			continue
		}
		// fmt.Printf("Title = #%v#\n", title) //判断左右两边是否存在空格的办法
		// fmt.Printf("content= #%v#\n", content)
		fileTitles = append(fileTitles, title) //追加内容
		fileContents = append(fileContents, content)
	}
	// fmt.Println("filetiles = ", fileTitles)
	// fmt.Println("fileContents = ", fileContents)

	//把内容写入到文件
	StoreToFile(i, fileTitles, fileContents)
	pagechan <- i //告诉后台第几页抓取完毕
}

//抓取子页面
func SpiderOnePage(onePageUrl string) (title, content string, err error) {
	//爬取单个页面的内容
	onepageresult, err1 := HttpGet(onePageUrl)
	if err1 != nil {
		err = err1
		return
	}

	//取标题<h1>标题</h1>　　只取第一个
	title_reg1 := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	if title_reg1 == nil {
		fmt.Println("SpiderOnePage regexp.MustCompile title reg err")
		return
	}

	titless := title_reg1.FindAllStringSubmatch(onepageresult, -1)
	// fmt.Println(titless[0][1])
	title = strings.TrimSpace(titless[0][1])

	//取笑话内容
	content_reg1 := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev"`)
	if content_reg1 == nil {
		fmt.Println("SpiderOnePage regexp.MustCompile content reg err")
		return
	}
	contentss := content_reg1.FindAllStringSubmatch(onepageresult, -1)
	// fmt.Println(contentss[0][1])
	content = strings.TrimSpace(contentss[0][1])
	// content = strings.Replace(content, "\t", "", -1) 替换所有的\t为空
	// content = strings.Replace(content, "\r", "", -1) 替换所有的\r为空
	// content = strings.Replace(content, "\n", "", -1) 替换所有的\n为空
	// content = strings.Replace(content, "<br />", "", -1) 替换所有的<br />为空
	return
}

//处理需要爬取的页面函数
func DoWork(start, end int) {
	//拼接需要爬取的url
	pagechan := make(chan int)

	for i := start; i <= end; i++ {
		go SpiderMainPage(i, pagechan)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d抓取完毕\n", <-pagechan)
	}

}

//主入口函数
func main() {
	var start, end int
	fmt.Println("请输入起始页：")
	fmt.Scan(&start)
	fmt.Println("请输入终止页：")
	fmt.Scan(&end)

	DoWork(start, end)
}
