package main

import (
	_ "ihome/models"
	_ "ihome/routers"

	"strings"

	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	//可以直接在main函数中进行部分函数的调试，比如调用controllers.upload函数
	//但是在main函数中，会不断的循环执行
	ignoreStaticPath()
	beego.Run()
	// beego.Run(":8899") //可以直接更换端口，优先于conf/app.conf中的配置
}

//所有的路径，都让TransparentStatic先处理一下
func ignoreStaticPath() {
	//设置fastdfs 图片的静态地址; 如果后台有fastdfs 进行数据存放，则存放fastdfs的目录。
	//也可以是直接使用普通的路径进行存放
	beego.SetStaticPath("group1/M00", "imagedata/stroage_data/data")
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

//对于传入的路径，先判断路径是否含有api字符串，有则跳过。无则加上static/html字段
func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	beego.Debug("request url:", orpath)
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)
}
