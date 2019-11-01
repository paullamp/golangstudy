package controllers

import (
	"encoding/json"
	"path"

	"github.com/weilaihui/fdfs_client"

	// "strings"

	//导入数据库的相关操作模块
	"ihome/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UsersController struct {
	beego.Controller
}

func (this *UsersController) RetData(data map[string]interface{}) {
	this.Data["json"] = data
	this.ServeJSON()
}

func (this *UsersController) GetUsers() {
	resp := make(map[string]interface{})

	//返回查询成功的内容
	resp["errno"] = models.RECODE_DBERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
	this.RetData(resp)
	return
}

/*
1. 设置路由
2. 添加user.go
3. 配置文件修改copyrequestbody = true
4. json.Unmarshal
5. 操作orm写入数据
*/
func (this *UsersController) Reg() {
	//获取前端的JSON数据
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	json.Unmarshal(this.Ctx.Input.RequestBody, &resp)
	// beego.Info("RequestBody 内容：", resp)
	//返回查询成功的内容
	// resp["errno"] = 4001
	// resp["errmsg"] = "查询失败"
	// this.RetData(resp)

	/*
		mobile: "111"
		password: "111"
		sms_code: "111"
	*/
	//插入数据
	o := orm.NewOrm()
	user := models.User{}
	user.Name = resp["mobile"].(string)
	user.Mobile = resp["mobile"].(string)
	user.Password_hash = resp["password"].(string)
	id, err := o.Insert(&user)
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	beego.Info("注册成功", id)
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	this.SetSession("name", user.Name)
	this.SetSession("user_id", user.Id)
	this.SetSession("mobile", user.Mobile)
	return
}

func (this *UsersController) Postavatar() {
	//定义最终需要上传的map变量，之后通过RetData返回json值
	resp := make(map[string]interface{})
	defer this.RetData(resp)

	//读取需要上传的文件
	filehandler, fileheader, err := this.GetFile("avatar")
	if err != nil {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		beego.Info("读取上传文件失败")
		return
	}

	//2. 若得文件后缀名
	suffix := path.Ext(fileheader.Filename)

	//3. 存储文件到fastdfs上
	//3.1 建立fdfs连接客户端
	fdfsClient, errfdfs := fdfs_client.NewFdfsClient("conf/client.conf")
	if errfdfs != nil {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		beego.Info("连接fdfs 服务器错误")
		return
	}
	//3.2 读filehander的内容读取到buffer中，然后执行上传
	fileBuffer := make([]byte, fileheader.Size)
	_, errFileHandlerRead := filehandler.Read(fileBuffer)
	if errFileHandlerRead != nil {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		beego.Info("读取filehandler内容出错")
		return
	}
	// suffix 读出来的是类似.jpg的内容，需要将.去掉，所以取suffix[1:] 切片
	upresp, errupload := fdfsClient.UploadByBuffer(fileBuffer, suffix[1:])
	if errupload != nil {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		beego.Info("上传至fdfs错误")
		return
	}
	// groupname := upresp.GroupName
	// fileid := upresp.RemoteFileId
	//4. 从session 里获取userid
	user_id := this.GetSession("user_id")

	//5. 从数据库中更新用户数据库中的内容
	o := orm.NewOrm()
	var user models.User
	qs := o.QueryTable("user")
	qs.Filter("Id", user_id).One(&user)
	user.Avatar_url = upresp.RemoteFileId
	_, errUpdate := o.Update(&user)

	if errUpdate != nil {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		beego.Info("更新数据库错误")
		return
	}
	urlMap := make(map[string]string)
	url := "http://10.1.89.154/" + upresp.RemoteFileId
	beego.Info("upresp.RemoteFileId [upresp.RemoteFileId的值是]:", upresp.RemoteFileId)
	beego.Info("Url fro avatar [拼接后返回的图片地址是]:", url)
	urlMap["avatar_url"] = url
	resp["data"] = urlMap
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	return

}

func (this *UsersController) GetUserData() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	// 1. 从session中获取user_id
	user_id := this.GetSession("user_id")
	// 2. 从数据库中拿到user_id 对应的user值
	user := models.User{Id: user_id.(int)}
	//3. 操作数据库，获取数据库信息
	o := orm.NewOrm()
	err := o.Read(&user)
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = user

}

func (this *UsersController) UpdateName() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)

	// 1.  从session 中取得user_id
	user_id := this.GetSession("user_id")

	//2. 获取前端传过来的数据
	UserName := make(map[string]string)
	json.Unmarshal(this.Ctx.Input.RequestBody, &UserName)
	beego.Info("get Username is :", UserName)

	//3. 更新数据库中的数据
	o := orm.NewOrm()
	user := models.User{Id: user_id.(int)}
	if err := o.Read(&user); err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		beego.Info("o.Read读取数据库的值失败")
		return
	}
	beego.Info("o.Read 读取的user 的值：", user)
	user.Name = UserName["name"]
	beego.Info("user 查询结果是：", user)
	_, err := o.Update(&user)
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return

	}
	//4. 更新session 中的数据
	this.SetSession("name", user.Name)

	//5.把数据打包返回给前端
	resp["data"] = UserName
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	return

}

func (this *UsersController) GetAuth() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)

	//1. 通过sessiona获取用户id
	user_id := this.GetSession("user_id")

	//2. 通过sessionid 查询数据库中的用户信息
	user := models.User{Id: user_id.(int)}
	o := orm.NewOrm()
	if err := o.Read(&user); err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	//3.将用户数据返回给客户端
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = user
	return
}

func (this *UsersController) PostAuth() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)

	// 1.获取用户的session_id
	user_id := this.GetSession("user_id")
	//2. 获取用户输入的名字及身份证号

	userinput := make(map[string]string)
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &userinput); err != nil {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}
	beego.Info("获取用户输入的信息：", userinput)

	//3. 更新数据库
	user := models.User{Id: user_id.(int)}
	o := orm.NewOrm()
	if err := o.Read(&user); err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	beego.Info("读取的数据库的用户信息：", user)
	user.Id_card = userinput["id_card"]
	user.Real_name = userinput["real_name"]
	//更新用户数据
	if _, err := o.Update(&user); err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	return
}
