package controllers

import (
	"encoding/json"
	"ihome/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type HouseController struct {
	beego.Controller
}

func (this *HouseController) RetData(data map[string]interface{}) {
	this.Data["json"] = data
	this.ServeJSON()
}

func (this *HouseController) GetHouseData() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	//1. 从session中获取user_id
	user_id := this.GetSession("user_id")

	//2. 从数据库中查询对应user_id的房子信息
	houses := []models.House{}
	o := orm.NewOrm()
	qs := o.QueryTable("house")
	num, err := qs.Filter("user__id", user_id.(int)).All(&houses)
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		beego.Info("查询房屋集合出错")
		return
	}
	if num == 0 {
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		beego.Info("查询数据库内无数据")
		return
	}
	respData := make(map[string]interface{})
	respData["houses"] = houses
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = respData
	return
}

func (this *HouseController) PostHouseData() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)

	//获取前端的输入请求数据
	reqData := make(map[string]interface{})
	json.Unmarshal(this.Ctx.Input.RequestBody, &reqData)
	beego.Info("获取房源的输入请求为：", reqData)
	house := models.House{}
	/*  将获取的map的信息，赋值给house结构体
	acreage: "121"
	address: "1212"
	area_id: "1"
	beds: "1212"
	capacity: "121"
	deposit: "1212"
	facility: ["1", "2", "11", "13"]
	max_days: "12"
	min_days: "121"
	price: "1212"
	room_count: "1212"
	title: "12312"
	unit: "1212"
	*/
	house.Title = reqData["title"].(string)
	price, _ := strconv.Atoi(reqData["price"].(string))

	house.Price = price
	house.Address = reqData["address"].(string)
	room_count, _ := strconv.Atoi(reqData["room_count"].(string))
	house.Room_count = room_count
	house.Unit = reqData["unit"].(string)
	house.Beds = reqData["beds"].(string)
	minDay, _ := strconv.Atoi(reqData["min_days"].(string))
	maxDay, _ := strconv.Atoi(reqData["max_days"].(string))
	house.Min_days = minDay
	house.Max_days = maxDay
	//获取sessionid 并且查询user的内容
	o := orm.NewOrm()
	user_id := this.GetSession("user_id")
	user := models.User{Id: user_id.(int)}
	if err := o.Read(&user); err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		beego.Info("通过user_id查询user数据：", err)
	}
	beego.Info("查询user_id对应的用户为：", user)
	house.User = &user
	faces := []models.Facility{}
	// for _, fid := range reqData["facility"].([]string) {
	for _, fid := range reqData["facility"].([]interface{}) {
		f_int_id, _ := strconv.Atoi(fid.(string))
		fac := models.Facility{Id: f_int_id}
		faces = append(faces, fac)
	}
	beego.Info("获取的faces信息：", faces)
	areaid, _ := strconv.Atoi(reqData["area_id"].(string))
	area := models.Area{Id: areaid}
	house.Area = &area

	//先切入以下house信息

	house_id, err := o.Insert(&house)
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		beego.Info("插入house表，获取id失败：", err)
		return
	}
	house.Id = int(house_id)

	m2m := o.QueryM2M(&house, "Facilities")
	num, errM2M := m2m.Add(faces)
	if errM2M != nil || num == 0 {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		beego.Info("M2M操作Facilities失败", errM2M)
		return
	}
	respData := make(map[string]interface{})
	respData["house_id"] = house_id
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = reqData
}

func (this *HouseController) GetSingleHouseData() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	//1. 从session中获取用户id
	user_id := this.GetSession("user_id").(int)
	beego.Info("获取到的user_id是：", user_id)
	//2. 获取用户输入信息
	house_id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	//3. 从缓存中读取用户信息

	//4. 从数据库中进行关联查询，查找对应的房屋信息
	o := orm.NewOrm()
	house := models.House{Id: house_id}
	if err := o.Read(&house); err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	respData := make(map[string]interface{})

	// rep2 := make(map[string]interface{})

	// //更新respData内的数据字段
	// respData["acreage"] = house.Acreage
	// respData["address"] = house.Address
	// respData["beds"] = house.Beds
	// respData["capacity"] = house.Capacity
	// // respData[""] = house.c/
	// respData["deposit"] = house.Deposit
	// respData["facilities"] = house.Facilities
	// respData["img_urls"] = house.Images
	// respData["max_days"] = house.Max_days
	// respData["min_days"] = house.Min_days
	// respData["price"] = house.Price
	// rep2["house"] = respData
	//关联查询
	o.LoadRelated(&house, "Area")
	o.LoadRelated(&house, "User")
	o.LoadRelated(&house, "Images")
	o.LoadRelated(&house, "Facilities")

	//关联读取用户信息
	user := models.User{Id: user_id}
	o.Read(&user)
	house.User = &user
	// for _, fac := range house.Facilities {
	// 	house.Facilities = append(house.Facilities, fac)
	// }
	beego.Info("House 设施内容如下：", house.Facilities)
	//5. 查询的数据结果存入缓存中
	respData["house"] = house
	//6. 返回正确的结果
	beego.Info("待返回的信息：", respData)
	resp["data"] = respData
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	return
}
