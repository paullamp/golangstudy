package controllers

import (
	"ihome/models"

	"github.com/astaxie/beego"
)

type TfastdfsControllerT struct {
	beego.Controller
}

func (this *TfastdfsControllerT) Tfastdfs() {
	var filename string = "a.txt"
	beego.Info("Filename is : ", filename)
	models.BeegoUploadByFilename(filename)
}

func (this *TfastdfsControllerT) IndexHello() {
	// this.Ctx.ResponseWriter("Helloworld")
	this.Ctx.WriteString("helloChina")
}
