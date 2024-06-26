package routers

import (
	"myapp/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	//new line
	beego.Router("/hello",&controllers.OwnTestController{} , "get:SayHello")
	beego.Router("/material", &controllers.OwnTestController{}, "get:GetMaterial")
}
