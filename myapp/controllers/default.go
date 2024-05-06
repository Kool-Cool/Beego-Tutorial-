package controllers

import (
	
	"myapp/models"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
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


func init() {
	orm.Debug = true // Enable debug logs
    orm.RegisterDriver("mysql", orm.DRMySQL)
    orm.RegisterDataBase("default", "mysql", "root:@/beegodb")
    orm.RegisterModel(new(models.Material))
	fmt.Println("Database initialized successfully")


}

func (mymate *OwnTestController) GetMaterial(){
	o := orm.NewOrm()
	materials := make([]*models.Material, 0)
	qs := o.QueryTable("material")
	_, err := qs.All(&materials)
	if err != nil {
		mymate.Data["error"] = err
	} else {
		mymate.Data["Materials"] = materials
	}
	mymate.TplName = "show.html"
}
