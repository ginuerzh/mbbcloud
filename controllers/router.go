// user
package controllers

import (
	"encoding/json"
	//"fmt"
	"github.com/ginuerzh/mbbcloud/models"
	"log"
	//"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
	"github.com/garyburd/redigo/redis"
	"github.com/ginuerzh/mbbcloud/errors"
	"strings"
	"time"
)

type RouterController struct {
	BaseController
}

func (this *RouterController) RouterList() {
	var routers models.RouterList

	if err := routers.FindAll(0, pageSize); err != nil {
		this.Data["json"] = this.response(nil, err)
		this.ServeJson()
		return
	}

	list := routers.Routers()

	for i, _ := range list {
		if time.Since(list[i].LastAccess.Value()) < 5*time.Minute {
			list[i].Online = true
		}
	}

	this.Data["json"] = this.response(list, nil)
	this.ServeJson()
}

func (this *RouterController) Router() {
	var router models.Router
	id := this.Ctx.Input.Params(":id")
	for {
		if err := router.FindOneBy("imei", id); err != nil {
			this.Data["json"] = this.response(nil, err)
			break
		}

		this.Data["json"] = this.response(&router, nil)
		break
	}

	this.ServeJson()
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

		router.LoginTime = models.NewJsonTime(time.Now())
		if err := router.Save(); err != nil {
			log.Println(err)
			this.Data["json"] = this.response(nil, err)
			break
		}

		token := this.uuid()

		c := RedisPool.Get()
		defer c.Close()
		c.Send("SET", NSRouter+token, router.Imei)
		c.Send("EXPIRE", NSRouter+token, ttl)
		c.Flush()
		c.Receive()

		r := map[string]string{"access_token": token}
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

		c := RedisPool.Get()
		defer c.Close()
		if s, err := redis.String(c.Do("GET", NSRouter+router.AccessToken)); err != nil {
			this.Data["json"] = this.response(nil, &errors.AuthError)
			break
		} else {
			router.Imei = s
			c.Do("EXPIRE", NSRouter+router.AccessToken, ttl)
		}

		if err := dbRouter.FindOneBy("imei", router.Imei); err != nil {
			this.Data["json"] = this.response(nil, err)
			break
		}

		dbRouter.CurDownlod = router.Download
		dbRouter.CurUpload = router.Upload
		dbRouter.LastAccess = models.NewJsonTime(time.Now())
		if router.Users != nil {
			dbRouter.Users = router.Users
		}

		if err := dbRouter.Save(); err != nil {
			log.Println(err)
			this.Data["json"] = this.response(nil, err)
			break
		}

		r := map[string]interface{}{"type": "info", "msg": "ok"}

		if s, err := redis.String(c.Do("RPOP", NSMQ+router.Imei)); err == nil {
			z := strings.SplitN(s, ":", 2)
			if len(z) == 2 {
				r["type"] = z[0]
				r["msg"] = z[1]
			}
		}

		this.Data["json"] = this.response(r, nil)
		break
	}

	this.ServeJson()
}
