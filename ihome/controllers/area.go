package controllers

import (
	"encoding/json"
	// "fmt"
	"time"

	"github.com/astaxie/beego/cache"

	// "encoding/json"
	//导入数据库的相关操作模块
	"ihome/models"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
)

type AreaController struct {
	beego.Controller
}

type Area struct {
	Errorid  int
	ErrorMsg string
}

func (c *AreaController) GetIndex() {
	c.Ctx.WriteString("Hello Getindex")
}
func (c *AreaController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
func (this *AreaController) RetJsonData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}
func (c *AreaController) GetArea1() {
	beego.Info("Connected to server")

	//定义resp map变量，用于存储需要返回的信息
	var resp map[string]interface{}
	resp = make(map[string]interface{})
	defer c.RetJsonData(resp)
	//1. 从数据库获取数据
	var areas []models.Area
	o := orm.NewOrm()
	_, err := o.QueryTable("area").All(&areas)
	if err != nil {
		beego.Error("查询数据库错误,", err)
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	beego.Info("areas infos:", areas) //此处将正常查询的结果打印到终端
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = areas
	//2. 打包成json并且返回

	//3.将json数据返回给前端

}

func (c *AreaController) GetArea() {
	//1. 从数据库获取信息
	var areas []models.Area
	resp := make(map[string]interface{})
	defer c.RetJsonData(resp)

	// 从redis缓存中取数据
	cache_conn, errNewCache := cache.NewCache("redis", `{"key":"ihome","conn":"10.0.197.17:6379","dbNum":"0","password":"mypass"}`)
	if errNewCache != nil {
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		beego.Info("create New cache error: ", errNewCache)
		return
	}
	// errPut := cache_conn.Put("aaa", "333", 60*60*time.Second)
	// if errPut != nil {
	// 	beego.Info("put cache to redis error:", errPut)
	// 	return
	// }
	// res := cache_conn.Get("aaa")
	// beego.Info("redis get 信息：", res)
	// fmt.Printf("redis get by fmt: %s", res)

	areaData := cache_conn.Get("areaData")
	if areaData != nil {
		// resp["data"] = areaData
		//后面需要仔细研究一下interfer 往其他数据转的功能
		var areainfo interface{}
		json.Unmarshal(areaData.([]byte), &areainfo)
		resp["data"] = areainfo
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		return
	} else {
		beego.Info("Error in areaData Get 查询redis areaData错误")
	}
	//创建查询句柄
	o := orm.NewOrm()
	num, err := o.QueryTable("area").All(&areas)
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	if num == 0 {
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		return
	}
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = areas

	//把数据转换成json格式，存入缓存
	bytebuf, errMarshal := json.Marshal(areas)
	if errMarshal != nil {
		beego.Info("Change to Json failed", errMarshal)
		return
	}
	errPut := cache_conn.Put("areaData", string(bytebuf), 60*60*time.Second)
	if errPut != nil {
		beego.Info("ErrPut 插入area 信息错误：", errPut)
		return
	}
	// fmt.Printf("data stored in redis ：%s", cache_conn.Get("aaa"))
}

func (this *AreaController) Jtestjson() {
	area := Area{10, "Hello"}
	this.Data["json"] = &area
	this.ServeJSON()
}
