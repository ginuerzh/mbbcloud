// app
package controllers

import (
	"encoding/json"
	"github.com/ginuerzh/mbbcloud/errors"
	"github.com/ginuerzh/mbbcloud/models"
	//"github.com/ginuerzh/weedo"
	"labix.org/v2/mgo/bson"
	"log"
)

const (
	CLIENT_IOS     = "ios"
	CLIENT_ANDROID = "android"
)

type AppController struct {
	BaseController
}

func (this *AppController) AppList() {
	var apps models.AppList

	client := this.GetString("client")

	if err := apps.FindAll(0, 0); err != nil {
		this.Data["json"] = this.response(nil, err)
		this.ServeJson()
		return
	}

	list := apps.Apps()
	for i, _ := range list {
		list[i].Icon = this.fileUrl(list[i].Icon)
		list[i].RUrl = this.fileUrl(list[i].RUrl)
		if client == CLIENT_IOS {
			list[i].CUrl = this.fileUrl(list[i].IUrl)
		}
		list[i].CUrl = this.fileUrl(list[i].AUrl)

		list[i].IUrl = ""
		list[i].AUrl = ""
	}
	this.Data["json"] = this.response(list, nil)
	this.ServeJson()
}

func (this *AppController) App() {
	var app models.App
	id := this.Ctx.Input.Param[":all"]

	if !bson.IsObjectIdHex(id) {
		this.Data["json"] = this.response(nil, &errors.InvalidParamsError)
		this.ServeJson()
		return
	}

	if err := app.FindOneBy("_id", bson.ObjectIdHex(id)); err != nil {
		this.Data["json"] = this.response(nil, err)
	} else {
		app.Icon = this.fileUrl(app.Icon)
		app.RUrl = this.fileUrl(app.RUrl)
		app.IUrl = this.fileUrl(app.IUrl)
		app.AUrl = this.fileUrl(app.AUrl)
		this.Data["json"] = this.response(&app, nil)
	}

	this.ServeJson()
}

func (this *AppController) Pub() {
	var app models.App

	for {
		//log.Println(string(this.Ctx.Input.RequestBody))
		if err := json.Unmarshal(this.Ctx.Input.RequestBody, &app); err != nil {
			log.Println(err)
			this.Data["json"] = this.response(nil, &errors.JsonError)
			break
		}

		app.PubTime = models.NewJsonTime(bson.Now())
		app.UpdateTime = app.PubTime
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
