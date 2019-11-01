package controllers

import (
	"encoding/json"
	"ihome/models"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
)

type SessionController struct {
	beego.Controller
}

func (this *SessionController) RetData(data map[string]interface{}) {
	this.Data["json"] = data
	this.ServeJSON()
}

/*
a> 增加一个session模块
b> 在项目里启用session(app.conf/ main.go)
c> 在注册结束后，增加session内容
d> 在首页判断session内是否有值

*/
func (this *SessionController) GetSessionInfo() {
	resp := make(map[string]interface{})
	user := models.User{}
	// user.Name = "zhangsan"
	//返回查询成功的内容
	resp["errno"] = models.RECODE_DBERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)

	name := this.GetSession("name")

	if name != nil {
		user.Name = name.(string)
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		resp["data"] = user
	}
	this.RetData(resp)
	return
}

/*
	session 的退出
*/

func (this *SessionController) DeleteSessionData() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	this.DelSession("name")
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}

func (this *SessionController) Login() {

	//1. 得到用户信息
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	json.Unmarshal(this.Ctx.Input.RequestBody, &resp)
	//用户输入信息为： map[mobile:111 password:111]
	beego.Info("用户输入信息为：", resp)

	//2. 判断用户输入是否合法
	if resp["mobile"] == nil || resp["password"] == nil {
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		beego.Info("func quit in 判断用户输入")
		return
	}
	//3. 从数据库查询信息
	o := orm.NewOrm()
	user := models.User{Mobile: resp["mobile"].(string)}
	err := o.Read(&user, "mobile")
	beego.Info("从数据库查询的结果为：", user)
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		beego.Info("数据库查询错误：", err)
		return
	}

	//判断密码是否错误
	if user.Password_hash != resp["password"].(string) {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	//4. 添加session
	this.SetSession("name", user.Name)
	this.SetSession("mobile", user.Mobile)
	this.SetSession("user_id", user.Id)
	//5. 返回json数据给前端
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}
