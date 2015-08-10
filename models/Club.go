package models

import (
	"github.com/astaxie/beego/orm"
	"log"
)

func (this *Club) Create() (int64, bool) {
	o := orm.NewOrm()
	o.Using("default")
	id, err := o.Insert(this)
	if err != nil {
		log.Println(err)
		return -1, false
	}
	return id, true
}

func (this *Club) GetAll() []Club {
	o := orm.NewOrm()
	o.Using("default")
	var clubs []Club
	_, err := o.QueryTable("Club").All(&clubs)
	if err != nil {
		return nil
	}
	return clubs
}

func (this *Club) GetById() bool {
	o := orm.NewOrm()
	o.Using("default")
	if status := o.QueryTable("club").Filter("id", this.Id).One(this); status != nil {
		return false
	}
	return true
}

func (this *Club) GetByTime(t int) []Club {
	var clubs []Club
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw("select * from club where lastmod>=?", t).QueryRows(&clubs)
	if err != nil {
		return nil
	}
	return clubs
}

func (this *Club) UpdateByTitle() bool {
	o := orm.NewOrm()
	o.Using("default")
	var club Club
	err := o.QueryTable("club").Filter("title", this.Title).One(&club)
	if err != nil {
		log.Println(err)
		return false
	}
	this.Id = club.Id
	o.Update(this)
	return true
}
