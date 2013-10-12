package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/ginuerzh/mbbcloud/controllers"
	//"github.com/ginuerzh/mbbcloud/models"
	//"icecar/filters"
	//"labix.org/v2/mgo"
	"log"
	"strconv"
	"strings"
)

func main() {
	beego.Router("/", &controllers.WebController{})
	beego.Router("/web/store", &controllers.WebController{}, "get:Store")
	beego.Router("/web/login", &controllers.WebController{}, "get:LoginGet;post:LoginPost")
	beego.Router("/web/logout", &controllers.WebController{}, "get:Logout")
	beego.Router("/web/apps", &controllers.WebController{}, "get,post:Apps")
	beego.Router("/web/app/pub", &controllers.WebController{}, "get:PubGet")
	beego.Router("/web/app/update/:all", &controllers.WebController{}, "get:UpdateApp")
	beego.Router("/web/app/:all", &controllers.WebController{}, "get:App")
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

	var adminFilter = func(ctx *context.Context) {
		//log.Println(ctx.Input.Url() + " in fileter")
		uid := ctx.GetCookie("uid")
		//log.Println("uid:" + uid)
		z := strings.Split(uid, ":")
		rank := controllers.UserRank0
		if len(z) == 2 {
			r, _ := strconv.ParseInt(z[1], 10, 32)
			rank = int(r)
		}
		if rank != controllers.UserRank1 {
			ctx.Redirect(302, "/web/login")
		}
	}

	var userFilter = func(ctx *context.Context) {
		c := ctx.GetCookie("uid")
		z := strings.Split(c, ":")
		uid := ""
		rank := controllers.UserRank0
		if len(z) == 2 {
			uid = z[0]
			r, _ := strconv.ParseInt(z[1], 10, 32)
			rank = int(r)
		}
		uri := ctx.Input.Uri()
		id := uri[strings.LastIndex(uri, "/")+1:]
		//log.Println(id, uid)
		if rank != controllers.UserRank1 && id != uid {
			ctx.Redirect(302, "/web/login")
		}
	}

	beego.AddFilter("/web/routers", "BeforRouter", adminFilter)
	beego.AddFilter("/web/router/:id", "BeforRouter", userFilter)
	beego.AddFilter("/web/app/update/:all", "BeforRouter", adminFilter)
	beego.AddFilter("/web/app/pub", "BeforRouter", adminFilter)
	beego.AddFilter("/file/upload", "BeforRouter", adminFilter)
	beego.AddFilter("/file/del/:all", "BeforRouter", adminFilter)

	beego.Debug("start server")
	beego.Run()
}
