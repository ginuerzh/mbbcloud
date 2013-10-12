package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"github.com/ginuerzh/mbbcloud/errors"
	"github.com/ginuerzh/mbbcloud/models"
	"github.com/nu7hatch/gouuid"
	"io"
	"strings"
	"time"
)

var (
	RedisPool *redis.Pool = redis.NewPool(dial, 3)
	pageSize  int         = 12
	ttl                   = 5 * time.Minute

	NSPrefix      = "mbbcloud:"
	NSRouters     = NSPrefix + "routers"
	NSRouter      = NSPrefix + "router:"
	NSRouterUsers = NSPrefix + "users:"
	NSRouterUser  = NSPrefix + "user:"
	NSMQ          = NSPrefix + "mq:"
)

const (
	CookieSecret = "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
	UserRank0    = 0 // rank for nobody
	UserRank1    = 1 // rank admin
	UserRank2    = 2 // rank common user
)

type BaseController struct {
	beego.Controller
}

func dial() (redis.Conn, error) {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (this *BaseController) fileUrl(fid string) string {
	if len(fid) == 0 {
		return ""
	}
	return "http://" + this.Ctx.Request.Host + "/file/" + strings.Join(strings.Split(fid, ","), "/")
}

func (this *BaseController) timeString(time time.Time) string {
	return time.Format(models.JsonTimeFormat)
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

func (this *BaseController) fileMd5(file io.Reader) string {
	h := md5.New()
	io.Copy(h, file)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (this *BaseController) response(r interface{}, e *errors.Error) map[string]interface{} {
	if e == nil {
		e = &errors.NoError
	}
	return map[string]interface{}{"r": e.ErrNo(), "err": e.ErrMsg(), "result": r}
}
