package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() { //用于处理对应url请求的get方法
	c.Ctx.WriteString("Hello world")
}

func (c *UserController) GetInfo() {
	c.Ctx.WriteString("getinfo response successful")
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("user input id =" + id)
}

func (this *UserController) GetNum() {
	id := this.Ctx.Input.Param(":id")
	this.Ctx.WriteString("get data from input: " + id)

}

//处理/download/*.* 类似的路由，　例如：/download/my.png；http://127.0.0.1:8080/download/efg.hello
//输出结果：filename= efg  fileext=hello
func (this *UserController) GetFilename() {
	filename := this.Ctx.Input.Param(":path")
	fileext := this.Ctx.Input.Param(":ext")
	this.Ctx.WriteString("filename= " + filename + "  fileext=" + fileext)
}

//处理以/down2/* 的正则匹配路由,例如：/down2/myfile.txt
//输出结果：splat= myfile.txt
func (this *UserController) GetAllinfo() {
	splat := this.Ctx.Input.Param(":splat")
	this.Ctx.WriteString("splat=" + splat)
}

func (this *UserController) PostData() {
	this.Ctx.WriteString("this is post function")
}
