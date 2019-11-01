package main

import (
	"helloworld/models"
	_ "helloworld/routers"

	// 测试orm使用增删改查，　在主函数中仅是测试操作
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// 测试orm使用增删改查，　在主函数中仅是测试操作
)

func insertOrder() {
	o := orm.NewOrm()
	order := models.User_order{}
	order.Order_data = "golang book"
	order.User = &models.User{Id: 1}

	_, err := o.Insert(&order)
	if err != nil {
		beego.Info("update user_order error=", err)
		return
	}

}

//测试数据表的插入
func insertUser() {
	o := orm.NewOrm()
	user := models.User{} //生成一个结构体，供操作表元素以及orm插入使用
	user.Name = "zhangsna"
	user.Age = 18
	id, err := o.Insert(&user) //插入至数据表中
	if err != nil {
		// fmt.Println("insert err : ", err)
		beego.Info("Insert error", err)
		return
	}
	beego.Info("Insert successfult ", id)
}

func queryOrder() {
	var orders []*models.User_order
	o := orm.NewOrm()
	qs := o.QueryTable("User_order")
	// _, err := qs.Filter("Id", 1).All(&orders) // 查询本表内的id
	_, err := qs.Filter("user__id", 1).All(&orders) // 查询本表内的user_id属性（实际不存在，是与user表中取到），中间有两具下划线
	if err != nil {
		beego.Error("Error in get all: ", err)
		return
	}
	// beego.Info(orders)
	// beego.Info(result)
	for _, order := range orders {
		beego.Info("query successful:", order)
	}
}
func main() {
	// insertUser() //调用函数，使用insert功能生效 , 若放在此位置，每执行一次main,都会调用一次
	// updateUser()
	// deleteUser()
	// queryUser()
	// insertOrder()
	// queryOrder()
	beego.SetStaticPath("download", "down") //设置当url为download的下载从down目录里进行资源查找。
	beego.Run()
}
