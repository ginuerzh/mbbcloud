package main

import (
	"github.com/astaxie/beego"
	"github.com/ginuerzh/mbbcloud/controllers"
	//"github.com/ginuerzh/mbbcloud/models"
	//"icecar/filters"
	//"labix.org/v2/mgo"
	"log"
)

func main() {
	beego.Router("/", &controllers.WebController{})
	beego.Router("/web/apps", &controllers.WebController{}, "get,post:Apps")
	beego.Router("/web/app/update/:all", &controllers.WebController{}, "get:UpdateApp")
	beego.Router("/web/app/:all", &controllers.WebController{}, "get:App")
	beego.Router("/web/pub", &controllers.WebController{}, "get:PubGet")
	beego.Router("/web/routers", &controllers.WebController{}, "get:Routers")
	beego.Router("/web/router/:id", &controllers.WebController{}, "get:Router")
	beego.Router("/web/router/send_msg", &controllers.WebController{}, "post:SendMessage")
	beego.Router("/apps", &controllers.AppController{}, "get,post:AppList")
	beego.Router("/app/pub", &controllers.AppController{}, "post:Pub")
	beego.Router("/app/update", &controllers.AppController{}, "post:Update")
	beego.Router("/app/del/:id", &controllers.AppController{}, "get:Delete")
	beego.Router("/app/:all", &controllers.AppController{}, "get:App")
	beego.Router("/routers", &controllers.RouterController{}, "get,post:RouterList")
	beego.Router("/router/:id", &controllers.RouterController{}, "get:Router")
	beego.Router("/login", &controllers.RouterController{}, "get,post:Login")
	beego.Router("/poll", &controllers.RouterController{}, "get,post:Poll")
	beego.Router("/file/upload", &controllers.FileController{}, "post:Upload")
	beego.Router("/file/del/:all", &controllers.FileController{}, "get:Delete")
	beego.Router("/file/:id([0-9]+)/:key([0-9a-f]+)", &controllers.FileController{}, "get:Download")

	beego.SetStaticPath("/images", "static/img")
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "static/js")

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	beego.Debug("start server")
	beego.Run()
}
