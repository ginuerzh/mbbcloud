package main

import (
	"github.com/astaxie/beego"
	"github.com/ginuerzh/mbbcloud/controllers"
	"github.com/ginuerzh/mbbcloud/models"
	//"icecar/filters"
	"labix.org/v2/mgo"
	"log"
)

func main() {
	beego.Router("/login", &controllers.RouterController{}, "get,post:Login")
	beego.Router("/poll", &controllers.RouterController{}, "get,post:Poll")
	beego.Router("/apps", &controllers.AppController{}, "get,post:AppList")
	beego.Router("/file/upload", &controllers.FileController{}, "post:Upload")
	beego.Router("/file/del/:all", &controllers.FileController{}, "get:Delete")
	beego.Router("/file/:all", &controllers.FileController{}, "get:Download")

	session, err := mgo.Dial(beego.AppConfig.String("mongourl"))
	if err != nil {
		panic(err)
		return
	}
	models.DB = session.DB(beego.AppConfig.String("mongodb"))
	defer session.Close()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	beego.Debug("start server")
	beego.Run()
}
