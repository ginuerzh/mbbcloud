// web
package controllers

import (
	"encoding/json"
	"github.com/ginuerzh/mbbcloud/errors"
	"github.com/ginuerzh/mbbcloud/models"
	//"github.com/ginuerzh/weedo"
	"log"
)

type WebController struct {
	BaseController
}

func (this *WebController) Get() {
	this.Layout = "base.html"
	this.TplNames = "store.html"
}

func (this *WebController) PubGet() {
	this.Layout = "base.html"
	this.TplNames = "pub.html"
}

func (this *WebController) PubPost() {
	var app models.App

	for {
		//log.Println(string(this.Ctx.Input.RequestBody))
		if err := json.Unmarshal(this.Ctx.Input.RequestBody, &app); err != nil {
			log.Println(err)
			this.Data["json"] = this.response(nil, &errors.JsonError)
			break
		}
		if err := app.Save(); err != nil {
			this.Data["json"] = this.response(nil, &errors.JsonError)
			break
		}

		r := map[string]string{"id": app.Id.String()}
		this.Data["json"] = this.response(r, nil)
		break
	}

	this.ServeJson()
}

func (this *WebController) SendMessage() {
	id := this.GetString("id")
	msgType := this.GetString("type")
	msg := this.GetString("msg")

	c := RedisPool.Get()
	defer c.Close()

	c.Do("LPUSH", models.NSMQ+id, msgType+":"+msg)

	this.Data["json"] = this.response(nil, nil)
	this.ServeJson()
}
