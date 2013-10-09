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
	beego.Router("/apps", &controllers.WebController{}, "get,post:Apps")
	beego.Router("/pub", &controllers.WebController{}, "get:PubGet")
	beego.Router("/routers", &controllers.WebController{}, "get:Routers")
	beego.Router("/router/send_msg", &controllers.WebController{}, "post:SendMessage")
	beego.Router("/login", &controllers.RouterController{}, "get,post:Login")
	beego.Router("/poll", &controllers.RouterController{}, "get,post:Poll")
	beego.Router("/router/list", &controllers.RouterController{}, "get,post:RouterList")
	beego.Router("/app/list", &controllers.AppController{}, "get,post:AppList")
	beego.Router("/app/pub", &controllers.AppController{}, "post:Pub")
	beego.Router("/file/upload", &controllers.FileController{}, "post:Upload")
	beego.Router("/file/del/:all", &controllers.FileController{}, "get:Delete")
	beego.Router("/file/:all", &controllers.FileController{}, "get:Download")

	beego.SetStaticPath("/images", "static/img")
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "static/js")

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	beego.Debug("start server")
	beego.Run()
}
