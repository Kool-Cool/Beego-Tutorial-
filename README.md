# Beego-Tutorial-
Step by Step guide for Beego Framework 


## Installing [Beego](https://beego.wiki/docs/intro/introduction/)
---

- ### S1] Check for `bee` version
  in terminal
 ```
 bee version
  ```

if result is like continue 
```go
 [D]  init global config instance failed. If you do not use this, just ignore it.  open conf/app.conf: The system cannot find the path specified.
______
| ___ \
| |_/ /  ___   ___
| ___ \ / _ \ / _ \
| |_/ /|  __/|  __/
\____/  \___| \___| v2.1.0
```

else continue with step S2

- ### S2] installing beego
  
```
go install github.com/beego/bee/v2@latest

bee version 
```

## Initializing Beego 
---

- S1]
```
bee new myapp
```

- S2] 
```
cd myapp
```
- S3]
```
go get myapp
```

- S4]
```
bee run
```

- S5] go to [http://:8080](http://localhost:8080/)
  
Result should be 

![](https://github.com/Kool-Cool/Beego-Tutorial-/blob/main/images/home.png)

## Hello World
---

- S1] Controllers

  Go to controllers > default.go

  Make new function using `MainController`
  ```go
  package controllers

  import (
  	beego "github.com/beego/beego/v2/server/web"
  )
  
  type MainController struct {
  	beego.Controller
  }
  
  func (c *MainController) Get() {
  	c.Data["Website"] = "beego.vip"
  	c.Data["Email"] = "astaxie@gmail.com"
  	c.TplName = "index.tpl"
  }
  
  // New Lines
  func (hello *MainController) SayHello(){
  	hello.TplName = "hello.html"
  }

  ```

- S2] Views

  Go to views and make `hello.html`

  ```html
  <!DOCTYPE html>
  <html lang="en">
  <head>
      <meta charset="UTF-8">
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <title>Hello World !</title>
  </head>
  <body>
      <h1>Hello World !</h1>
      <p>We are using Beego Framwork for api rounting</p>
      
  </body>
  </html>

  ```

- S3] Routes
  Go to routers > router.go

  assign new route

  ```go
  package routers
  
  import (
  	"myapp/controllers"
  	beego "github.com/beego/beego/v2/server/web"
  )
  
  func init() {
      beego.Router("/", &controllers.MainController{})
  	  //new line
  	  beego.Router("/hello ",&controllers.MainController{} , "get:SayHello") 
  }


  ```

  Result should be 

  ![](https://github.com/Kool-Cool/Beego-Tutorial-/blob/main/images/HelloWorld.png)


## Making Own New Controller
---

NOTE:- [Make First Letter Capital](https://www.golinuxcloud.com/golang-variable-naming-convention/#:~:text=The%20convention%20also%20dictates%20that%20the%20names%20of%20global%20variables%20should%20start%20with%20uppercase%20letters%20if%20they%20are%20to%20be%20exported%2C%20meaning%20they%20will%20be%20accessible%20outside%20of%20the%20package.)

- S1] changes in `default.go`

  Make new structures
  ```go
  type OwnTestController struct{
  	beego.Controller
  }
  
  func (hello *OwnTestController) SayHello(){
  	hello.TplName = "hello.html"
  }
  ```

- S2] changes in `router.go`
  update the name of Controller

  ```go
  package routers
  
  import (
  	"myapp/controllers"
  	beego "github.com/beego/beego/v2/server/web"
  )
  
  func init() {
      beego.Router("/", &controllers.MainController{})
  	//new line
  	beego.Router("/hello",&controllers.OwnTestController{} , "get:SayHello")
  }

  ```


## Sending Data  
---

Use `funcName.Data['attribute'] = value`

```go

func (hello *OwnTestController) SayHello(){
	hello.Data["name"] = "This is my name !" 
	hello.TplName = "hello.html"
}

```

to access value , change html file `{{.attribute}}`

```html
<h2>{{.name}}</h2>
```

## Connecting DataBase (MySQL) in GOlang  
---
Please refer the [db.go](https://github.com/Kool-Cool/Beego-Tutorial-/blob/main/db.go)


## Connecting the DataBase with ORM
---

- S1] Make model

  Make new file under `models/material.go`

  ```go
	type Material struct {
		MaterialID   int    `orm:"column(material_id);pk"`
		MaterialName string `orm:"column(material_name)"`
		MaterialSubtype string `orm:"column(material_subtype)"`
	}

  ```


- S2] Establishing the connection with database

  Make new function `init()` to Establish connection with DataBase and Models
  at `controllers/default.go`

  ```go
	func init() {
		orm.Debug = true // Enable debug logs
	    orm.RegisterDriver("mysql", orm.DRMySQL)
							//"databaseType" , "userName:password@/NameOFDataBase"
	    orm.RegisterDataBase("default", "mysql", "root:@/beegodb")
	    orm.RegisterModel(new(models.Material))
		fmt.Println("Database initialized successfully")
	
	}
  ```

- S3] Using data from Model

    make new function to handle requests at `controllers/default.go`

    ```go
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
		// mymate.TplName = "show.html"
		mymate.Ctx.JSONResp(materials) 
	
	}
    ```

    change query as per required !
    
- S4] Updating The Routes at `router.go/router.go`

     ```
	beego.Router("/material", &controllers.OwnTestController{}, "get:GetMaterial"
     ```

	OutPut Result should be 

  	![](https://github.com/Kool-Cool/Beego-Tutorial-/blob/main/images/output.png)
  	
  
