package routers

import (
	"cmritfeed/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/event/?:time", &controllers.EventController{}, "*:Events")
	beego.Router("/event/create", &controllers.EventController{}, "*:Create")
	beego.Router("/event/update", &controllers.EventController{}, "*:Update")
	beego.Router("/attend/", &controllers.EventController{}, "*:Attend")
	beego.Router("/attending/:id", &controllers.EventController{}, "*:Attending")
	beego.Router("/club/?:time", &controllers.ClubController{}, "*:Clubs")
	beego.Router("/club/update/", &controllers.ClubController{}, "*:UpdateClub")
	beego.Router("/club/create/", &controllers.ClubController{}, "*:CreateClub")

}
