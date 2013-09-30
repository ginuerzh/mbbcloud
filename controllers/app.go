// app
package controllers

import (
	"encoding/json"
	"github.com/ginuerzh/mbbcloud/errors"
	"github.com/ginuerzh/mbbcloud/models"
	//"github.com/ginuerzh/weedo"
	"log"
)

type AppController struct {
	BaseController
}

func (this *AppController) AppList() {
	var apps models.AppList

	//client := this.GetString("client")

	if err := apps.FindAll(0, 0); err != nil {
		this.Data["json"] = this.response(nil, err)
		this.ServeJson()
		return
	}

	list := apps.Apps()
	for i, _ := range list {
		//list[i].Icon, _ = weedo.GetUrl(list[i].Icon)
		//list[i].RUrl, _ = weedo.GetUrl(list[i].RUrl)
		//list[i].IUrl, _ = weedo.GetUrl(list[i].IUrl)
		uri := "http://localhost:12345/file/"
		list[i].Icon = uri + list[i].Icon
		list[i].RUrl = uri + list[i].RUrl
		list[i].IUrl = uri + list[i].IUrl
	}
	this.Data["json"] = this.response(list, nil)
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
