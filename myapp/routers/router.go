package routers

import (
	"myapp/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/hello ",&controllers.MainController{} , "get:SayHello") // New Line
}
