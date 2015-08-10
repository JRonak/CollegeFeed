package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Club struct {
	Id          int `orm:"pk"`
	Title       string
	Description string
	President   string
	Contact     string
	Lastmod     int
}

type Event struct {
	Id          int `orm:"pk"`
	Title       string
	Description string
	Attendcount int
	Time        int
	Club        string
	Lastmod     int
}

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "ronak:ronak@/cmritfeed")
	orm.RegisterModel(new(Club))
	orm.RegisterModel(new(Event))
}
