package models

import (
	"github.com/astaxie/beego/orm"
	"log"
)

func (this *Event) Create() (int64, bool) {
	o := orm.NewOrm()
	o.Using("default")
	id, err := o.Insert(this)
	if err != nil {
		log.Println(err)
		return 0, false
	}
	return id, true
}

func (this *Event) UpdateByTitle() (int64, bool) {
	o := orm.NewOrm()
	o.Using("default")
	var event Event
	err := o.QueryTable("event").Filter("title", this.Title).One(&event)
	if err != nil {
		log.Println(err)
		return 0, false
	}
	this.Id = event.Id
	id, err := o.Update(this)
	if err != nil {
		log.Println(err)
		return 0, false
	}
	return id, true
}

func (this *Event) Delete() bool {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Delete(this)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (this *Event) GetAll() []Event {
	var events []Event
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.QueryTable("Event").All(&events)
	if err != nil {
		log.Println(err)
		return nil
	}
	return events
}

func (this *Event) GetById() bool {
	o := orm.NewOrm()
	o.Using("default")
	if i := o.Read(this); i != nil {
		return false
	}
	return true
}

func (this *Event) GetNew(id int) []Event {
	var events []Event
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw("select * from event where id>?", id).QueryRows(&events)
	if err != nil {
		return nil
	}
	return events
}

func (this *Event) GetByClub(club string) []Event {
	var events []Event
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw("select * from event where club=?", club).QueryRows(&events)
	if err != nil {
		return nil
	}
	return events
}

func (this *Event) GetAttendByTime(t int) []Event {
	var events []Event
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw("select * from event where time>=?", t).QueryRows(&events)
	if err != nil {
		return nil
	}
	return events
}

func (this *Event) GetAfter(t int) []Event {
	var events []Event
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw("select * from event where lastmod>=?", t).QueryRows(&events)
	if err != nil {
		return nil
	}
	return events
}

func (this *Event) AddAttend() bool {
	o := orm.NewOrm()
	o.Using("default")
	o.Read(this)
	this.Attendcount += 1
	_, err := o.Update(this)
	if err != nil {
		return false
	}
	return true
}
