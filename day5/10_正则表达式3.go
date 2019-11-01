package main

import (
	"fmt"
	"regexp"
)

func main() {
	message := `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<title>Go语言标准库文档中文版 | Go语言中文网 | Golang中文社区 | Golang中国</title>
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
	<meta http-equiv="X-UA-Compatible" content="IE=edge, chrome=1">
	<meta charset="utf-8">
	<link rel="shortcut icon" href="/static/img/go.ico">
	<link rel="apple-touch-icon" type="image/png" href="/static/img/logo2.png">
	<meta name="author" content="polaris <polaris@studygolang.com>">
	<meta name="keywords" content="中文, 文档, 标准库, Go语言,Golang,Go社区,Go中文社区,Golang中文社区,Go语言社区,Go语言学习,学习Go语言,Go语言学习园地,Golang 中国,Golang中国,Golang China, Go语言论坛, Go语言中文网">
	<meta name="description" content="Go语言文档中文版，Go语言中文网，中国 Golang 社区，Go语言学习园地，致力于构建完善的 Golang 中文社区，Go语言爱好者的学习家园。分享 Go 语言知识，交流使用经验">
</head>
<frameset cols="15,85">
	<frame src="/static/pkgdoc/i.html">
	<frame name="main" src="/static/pkgdoc/main.html" tppabs="main.html" >
	<noframes>
	</noframes>
</frameset>
	<div>哈哈</div>
	<div>GOODJOB</div>
	<div>你过来呀</div>
	<div>哈哈，可爱的世界</div>
</html>
	`

	//定义需要匹配的字符
	// reg1 := regexp.MustCompile(`<div>(.*)</div>`)
	reg1 := regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	// result := reg1.FindAllStringSubmatch(message, -1)
	result := reg1.FindAllStringSubmatch(message, -1)
	// fmt.Println(result)

	for _, text := range result {
		// fmt.Println("text[0]=", text[0]) //带有<div></div>的内容
		fmt.Println("text[1]", text[1]) //不带<div></div>的内容，即具体匹配的内容
	}
}
