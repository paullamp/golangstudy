package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//用户的数据库映射
type User struct {
	Id         int
	Name       string `orm:size(100)`
	Age        int
	User_order []*User_order `orm:"reverse(many)"` //一对多
}

//订单表的数据库映射
type User_order struct {
	Id         int
	Order_data string `orm:size(100)`
	User       *User  `orm:"rel(fk)"` //其实就是指定了一个foreign key
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:golangmysql@tcp(10.0.197.17:3306)/orm_test?charset=utf8")
	orm.RegisterModel(new(User), new(User_order))
	orm.RunSyncdb("default", false, true)
}
