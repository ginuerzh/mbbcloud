// web
package controllers

import (
	"github.com/ginuerzh/mbbcloud/errors"
	"github.com/ginuerzh/mbbcloud/models"
	//"github.com/ginuerzh/weedo"
	"fmt"
	"log"
)

type WebController struct {
	BaseController
}

func (this *WebController) Get() {
	this.Redirect("/web/store", 302)
}

func (this *WebController) Store() {
	this.Layout = "base.html"
	this.TplNames = "store.html"
}

func (this *WebController) LoginGet() {
	this.Layout = "base.html"
	this.TplNames = "login.html"
}

func (this *WebController) LoginPost() {
	var user models.User
	var router models.Router

	username := this.GetString("username")
	password := this.GetString("password")

	if err := user.FindOneBy("username", username); err != nil {
		if err := router.FindOneBy("imei", username); err != nil {
			this.Data["json"] = this.response(nil, &errors.UserNotFoundError)
			this.ServeJson()
			return
		} else {
			user.Username = username
			user.Rank = UserRank2
			user.Password = this.md5("123456")
			user.Save()
		}
	}
	if user.Password != this.md5(password) {
		this.Data["json"] = this.response(nil, &errors.PasswordError)
		this.ServeJson()
		return
	}

	uid := fmt.Sprintf("%s:%d", user.Username, user.Rank)
	this.Ctx.SetCookie("uid", uid, 3600, "/")
	this.Redirect("/", 302)
}

func (this *WebController) Logout() {
	this.Ctx.SetCookie("uid", "", -1, "/")
	this.Redirect("/", 302)
}

func (this *WebController) PubGet() {
	this.Layout = "base.html"
	this.TplNames = "pub.html"
}

func (this *WebController) Apps() {
	this.Layout = "base.html"
	this.TplNames = "store.html"
}

func (this *WebController) App() {
	this.Layout = "base.html"
	this.TplNames = "app.html"
}
func (this *WebController) UpdateApp() {
	this.Layout = "base.html"
	this.TplNames = "update.html"
}
func (this *WebController) Routers() {
	this.Layout = "base.html"
	this.TplNames = "router.html"
}
func (this *WebController) Router() {
	this.Layout = "base.html"
	this.TplNames = "router_info.html"
}
func (this *WebController) SendMessage() {
	id := this.GetString("id")
	msgType := this.GetString("type")
	msg := this.GetString("msg")

	c := RedisPool.Get()
	defer c.Close()

	c.Do("LPUSH", NSMQ+id, msgType+":"+msg)

	log.Printf("Send message %s(%s) to %s\n", msg, msgType, id)
	this.Data["json"] = this.response(nil, nil)
	this.ServeJson()
}
