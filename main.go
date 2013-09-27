package main

import (
	"github.com/astaxie/beego"
	"github.com/ginuerzh/mbbcloud/controllers"
	"github.com/ginuerzh/mbbcloud/models"
	//"icecar/filters"
	"github.com/garyburd/redigo/redis"
	"labix.org/v2/mgo"
	"log"
)

func main() {
	beego.Router("/login", &controllers.UserController{}, "get,post:Login")
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

	models.Redis, err = redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
		return
	}
	defer models.Redis.Close()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	beego.Debug("start server")
	beego.Run()
}
