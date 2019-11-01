package controllers

import (
	//导入数据库的相关操作模块
	"ihome/models"

	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
)

type HousesIndexController struct {
	beego.Controller
}

func (this *HousesIndexController) RetData(data map[string]interface{}) {
	this.Data["json"] = data
	this.ServeJSON()
}

func (this *HousesIndexController) GetHousesIndex() {
	resp := make(map[string]interface{})

	//返回查询成功的内容
	resp["errno"] = models.RECODE_NODATA
	resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
	this.RetData(resp)
	return
}
