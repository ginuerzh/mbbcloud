// user
package models

import (
	"github.com/ginuerzh/mbbcloud/errors"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	//"log"
	"time"
)

type User struct {
	Device   string
	Mac      string
	Ip       string
	Duration int
}

type Router struct {
	Imei        string `json:"id"`
	Mac         string
	Ip          string
	Download    uint64
	Upload      uint64
	CurDownlod  uint64 `bson:",omitempty"`
	CurUpload   uint64 `bson:",omitempty"`
	Online      bool
	LoginTime   time.Time `bson:"login_time"`
	LastAccess  time.Time `bson:"last_access,omitempty"`
	AccessToken string    `json:"access_token" bson:"access_token,omitempty"`
	Users       []User    `bson:",omitempty"`
}

func (this *Router) FindBy(key string, value interface{}) (e *errors.Error) {
	c := DB.C(C_Router)

	if err := c.Find(bson.M{key: value}).One(this); err != nil {
		if err == mgo.ErrNotFound {
			e = &errors.UserNotFoundError
		} else {
			e = &errors.DbError
		}
	}

	return
}

func (this *Router) Exists() (b bool) {
	c := DB.C(C_Router)

	count, _ := c.Find(bson.M{"imei": this.Imei}).Count()

	if count > 0 {
		b = true
	}

	return
}

func (this *Router) Save() *errors.Error {
	c := DB.C(C_Router)

	if len(this.Imei) == 0 || len(this.Mac) == 0 || len(this.Ip) == 0 {
		return &errors.InvalidParamsError
	}

	if _, err := c.Upsert(bson.M{"imei": this.Imei}, this); err != nil {
		return &errors.DbError
	}

	return nil
}

func (this *Router) SetOnline(online bool) (e *errors.Error) {
	c := DB.C(C_Router)

	if online && len(this.AccessToken) == 0 {
		e = &errors.AccessError
		return
	}

	change := bson.M{
		"$set": bson.M{
			"online":       online,
			"access_token": this.AccessToken,
		},
	}

	if err := c.Update(bson.M{"imei": this.Imei}, change); err != nil {
		e = &errors.DbError
		return
	}

	this.Online = online
	return
}
