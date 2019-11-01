package models

import (
	"github.com/astaxie/beego"
	// 测试orm使用增删改查，　在主函数中仅是测试操作
	"github.com/astaxie/beego/orm"
)

//测试数据表的插入
func insertUser() {
	o := orm.NewOrm()
	user := User{} //生成一个结构体，供操作表元素以及orm插入使用
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

//测度数据表的查询
func queryUser() {
	o := orm.NewOrm()
	user := User{Id: 1} //创建一个有一个元素初始化了的user结构体，后期传入至orm使用
	err := o.Read(&user)
	if err != nil {
		beego.Info("query read error:", err)
		return
	}
	beego.Info("query successful: ", user)
}

//测试数据表的更新
func updateUser() {
	o := orm.NewOrm()
	user := User{Id: 1, Name: "wangxx", Age: 25}
	o.Update(&user)
	beego.Info("updated User is :", user) //因为传递的是指针引用，可以将值直接在user中修改
}

//测试数据表的删除
func deleteUser() {
	o := orm.NewOrm()
	user := User{Id: 1}
	_, err := o.Delete(&user)
	if err != nil {
		beego.Info("delete failed ", err)
		return
	}
}
