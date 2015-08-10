package controllers

import (
	"cmritfeed/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type EventController struct {
	beego.Controller
}

var (
	failjson string
	passjson string
)

func init() {
	status := make(map[string]bool)
	status["status"] = true
	b, _ := json.Marshal(&status)
	passjson = string(b)
	status["status"] = false
	b, _ = json.Marshal(&status)
	failjson = string(b)
}

func (this *EventController) Create() {
	password := this.GetString("password")
	if password != "ronak123" {
		this.Data["json"] = failjson
		this.ServeJson()
		return
	}
	title := this.GetString("title")
	description := this.GetString("description")
	t := this.GetString("time")
	lastmod := int(time.Now().Unix())
	event := models.Event{}
	event.Title = title
	event.Description = description
	tim, _ := strconv.Atoi(t)
	event.Time = tim
	event.Lastmod = lastmod
	_, status := event.Create()
	if status != true {
		this.Data["json"] = failjson
		this.ServeJson()
		return
	} else {
		this.Data["json"] = passjson
		this.ServeJson()
		return
	}
}

func (this *EventController) Update() {
	password := this.GetString("password")
	if password != "ronak123" {
		this.Data["json"] = failjson
		this.ServeJson()
		return
	}
	title := this.GetString("title")
	description := this.GetString("description")
	t := this.GetString("time")
	lastmod := int(time.Now().Unix())
	event := models.Event{}
	event.Title = title
	event.Description = description
	tim, _ := strconv.Atoi(t)
	event.Time = tim
	event.Lastmod = lastmod
	_, status := event.UpdateByTitle()
	if status != true {
		this.Data["json"] = failjson
		this.ServeJson()
		return
	} else {
		this.Data["json"] = passjson
		this.ServeJson()
		return
	}
}
func (this *EventController) Events() {
	time, _ := strconv.Atoi(this.Ctx.Input.Param(":time"))
	t := models.Event{}
	ts := t.GetAfter(time)
	b, _ := json.Marshal(&ts)
	this.Data["json"] = string(b)
	this.ServeJson()
}

func (this *EventController) Attend() {
	t := models.Event{}
	timenow := time.Now().Unix()
	attend := t.GetAttendByTime(int(timenow))
	for i, _ := range attend {
		attend[i].Club = ""
		attend[i].Description = ""
		attend[i].Title = ""
	}
	b, _ := json.Marshal(&attend)
	this.Data["json"] = string(b)
	this.ServeJson()
}

func (this *EventController) Attending() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	t := models.Event{}
	t.Id = id
	t.AddAttend()
	this.ServeJson()
}
