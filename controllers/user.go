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
)

type UserController struct {
	BaseController
}

func (this *UserController) Login() {
	var user models.User

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &user); err != nil {
		log.Println(err)
		this.Data["json"] = this.response(nil, &errors.JsonError)
		this.ServeJson()
		return
	}

	if err := user.Save(); err != nil {
		log.Println(err)
		this.Data["json"] = this.response(nil, err)
		this.ServeJson()
		return
	}
	token := this.uuid()
	r := map[string]string{"access_token": token}
	this.Data["json"] = this.response(r, nil)

	this.ServeJson()
}
