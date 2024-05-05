package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}



func (c *MainController) Get() {
	// c.Data["Website"] = "beego.vip"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"
	// c.Ctx.WriteString("hello world 🤓 \n Just text (not html file)")
	c.TplName = "home.html"
}

// New Lines
type OwnTestController struct{
	beego.Controller
}

func (hello *OwnTestController) SayHello(){
	hello.Data["name"] = "This is my name !" 
	hello.TplName = "hello.html"
}
