package controllers

import (
	"ihome/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type OrderController struct {
	beego.Controller
}

func (this *OrderController) RetData(data map[string]interface{}) {
	this.Data["json"] = data
	this.ServeJSON()
}

func (this *OrderController) GetUserOrder() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)

	//1. 获取user_id
	user_id := this.GetSession("user_id").(int)

	//2. 根据url 获取当前url操作的角色
	role := this.GetString("role")
	beego.Info("当前的角色是role 是：", role)
	if role == "custom" {
		//租客处理逻辑
		orders := []models.OrderHouse{}
		o := orm.NewOrm()
		qs := o.QueryTable("OrderHouse")
		qs.Filter("user__id", user_id).All(&orders)
		user := models.User{Id: user_id}
		for _, order := range orders {
			order.User = &user
			o.LoadRelated(order, "User")
		}
		respData := make(map[string]interface{})
		respData["order"] = orders
		resp["data"] = respData
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		return
	}

	if role == "landlord" {

	}

	if role == "" {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
	}
}
