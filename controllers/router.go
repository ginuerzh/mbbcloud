// user
package controllers

import (
	"encoding/json"
	//"fmt"
	"github.com/ginuerzh/mbbcloud/models"
	"log"
	//"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
	"github.com/ginuerzh/mbbcloud/errors"
	"time"
)

type RouterController struct {
	BaseController
}

func (this *RouterController) Login() {
	var router models.Router
	//log.Println(string(this.Ctx.Input.RequestBody))

	for {
		if err := json.Unmarshal(this.Ctx.Input.RequestBody, &router); err != nil {
			log.Println(err)
			this.Data["json"] = this.response(nil, &errors.JsonError)
			break
		}

		router.LoginTime = time.Now()
		if err := router.Save(); err != nil {
			log.Println(err)
			this.Data["json"] = this.response(nil, err)
			break
		}

		router.AccessToken = this.uuid()

		if err := router.SetOnline(true); err != nil {
			this.Data["json"] = this.response(nil, err)
			break
		}

		r := map[string]string{"access_token": router.AccessToken}
		this.Data["json"] = this.response(r, nil)

		break
	}

	this.ServeJson()
}

func (this *RouterController) Poll() {
	var router, dbRouter models.Router

	for {
		if err := json.Unmarshal(this.Ctx.Input.RequestBody, &router); err != nil {
			log.Println(err)
			this.Data["json"] = this.response(nil, &errors.JsonError)
			break
		}
		log.Println(router.AccessToken)
		if err := dbRouter.FindOneBy("access_token", router.AccessToken); err != nil {
			this.Data["json"] = this.response(nil, err)
			break
		}

		dbRouter.CurDownlod = router.Download
		dbRouter.CurUpload = router.Upload
		dbRouter.LastAccess = time.Now()
		if router.Users != nil {
			dbRouter.Users = router.Users
		}

		if err := dbRouter.Save(); err != nil {
			log.Println(err)
			this.Data["json"] = this.response(nil, err)
			break
		}

		r := map[string]interface{}{"type": "info", "msg": "ok"}
		this.Data["json"] = this.response(r, nil)
		break
	}

	this.ServeJson()
}
