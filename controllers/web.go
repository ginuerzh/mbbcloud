// web
package controllers

import (
	//"github.com/ginuerzh/mbbcloud/errors"
	//"github.com/ginuerzh/mbbcloud/models"
	//"github.com/ginuerzh/weedo"
	"log"
)

type WebController struct {
	BaseController
}

func (this *WebController) Get() {
	log.Println(this.Ctx.Request.RemoteAddr)
	this.Layout = "base.html"
	this.TplNames = "store.html"
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

	this.Data["json"] = this.response(nil, nil)
	this.ServeJson()
}
