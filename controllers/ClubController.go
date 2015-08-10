package controllers

import (
	"cmritfeed/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type ClubController struct {
	beego.Controller
}

func (this *ClubController) Clubs() {
	time, _ := strconv.Atoi(this.Ctx.Input.Param(":time"))
	c := models.Club{}
	clubs := c.GetByTime(time)
	b, _ := json.Marshal(&clubs)
	this.Data["json"] = string(b)
	this.ServeJson()
}

func (this *ClubController) CreateClub() {
	password := this.GetString("password")
	if password != "ronak123" {
		this.Data["json"] = failjson
		this.ServeJson()
		return
	}
	title := this.GetString("title")
	description := this.GetString("description")
	president := this.GetString("president")
	contact := this.GetString("contact")
	club := models.Club{}
	club.President = president
	club.Contact = contact
	club.Title = title
	club.Description = description
	club.Lastmod = int(time.Now().Unix())
	_, err := club.Create()
	if err != true {
		this.Data["json"] = failjson
		this.ServeJson()
		return
	} else {
		this.Data["json"] = passjson
		this.ServeJson()
		return
	}
}

func (this *ClubController) UpdateClub() {
	password := this.GetString("password")
	if password != "ronak123" {
		this.Data["json"] = failjson
		this.ServeJson()
		return
	}
	title := this.GetString("title")
	description := this.GetString("description")
	president := this.GetString("president")
	contact := this.GetString("contact")
	club := models.Club{}
	club.President = president
	club.Contact = contact
	club.Title = title
	club.Description = description
	club.Lastmod = int(time.Now().Unix())
	err := club.UpdateByTitle()
	if err != true {
		this.Data["json"] = failjson
		this.ServeJson()
		return
	} else {
		this.Data["json"] = passjson
		this.ServeJson()
		return
	}
}
