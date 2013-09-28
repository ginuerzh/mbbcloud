package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"github.com/ginuerzh/mbbcloud/errors"
	"github.com/nu7hatch/gouuid"
	"io"
)

func dial() (redis.Conn, error) {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		return nil, err
	}

	return c, nil
}

var (
	Pool *redis.Pool = redis.NewPool(dial, 3)
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) uuid() string {
	u4, err := uuid.NewV4()
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}

	return u4.String()
}

func (this *BaseController) md5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (this *BaseController) response(r interface{}, e *errors.Error) map[string]interface{} {
	if e == nil {
		e = &errors.NoError
	}
	return map[string]interface{}{"r": e.ErrNo(), "err": e.ErrMsg(), "result": r}
}
