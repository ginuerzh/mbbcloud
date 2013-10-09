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
	beego.Router("/apps", &controllers.AppController{}, "get,post:AppList")
	beego.Router("/app/pub", &controllers.AppController{}, "post:Pub")
	beego.Router("/web/app/:all", &controllers.WebController{}, "get:App")
	beego.Router("/web/pub", &controllers.WebController{}, "get:PubGet")
	beego.Router("/web/routers", &controllers.WebController{}, "get:Routers")
	beego.Router("/web/router/send_msg", &controllers.WebController{}, "post:SendMessage")
	beego.Router("/router/list", &controllers.RouterController{}, "get,post:RouterList")
	beego.Router("/login", &controllers.RouterController{}, "get,post:Login")
	beego.Router("/poll", &controllers.RouterController{}, "get,post:Poll")
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
