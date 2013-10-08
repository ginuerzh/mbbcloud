// app
package controllers

import (
	//"encoding/json"
	//"github.com/ginuerzh/mbbcloud/errors"
	"github.com/ginuerzh/mbbcloud/models"
	//"github.com/ginuerzh/weedo"
	//"log"
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
		uri := this.Ctx.Request.Host + "/file/"
		if len(list[i].Icon) > 0 {
			list[i].Icon = uri + list[i].Icon
		}
		if len(list[i].RUrl) > 0 {
			list[i].RUrl = uri + list[i].RUrl
		}
		if len(list[i].IUrl) > 0 {
			list[i].IUrl = uri + list[i].IUrl
		}
	}
	this.Data["json"] = this.response(list, nil)
	this.ServeJson()
}
