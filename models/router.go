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

func (this *Router) FindOneBy(key string, value interface{}) (e *errors.Error) {
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{key: value}).One(this)
	}

	if err := withCollection(C_Router, query); err != nil {
		if err == mgo.ErrNotFound {
			e = &errors.UserNotFoundError
		} else {
			e = &errors.DbError
		}
	}

	return
}

func (this *Router) Exists() bool {
	count := 0
	query := func(c *mgo.Collection) error {
		var err error
		count, err = c.Find(bson.M{"imei": this.Imei}).Count()
		return err
	}

	if withCollection(C_Router, query); count > 0 {
		return true
	}

	return false
}

func (this *Router) Save() *errors.Error {
	if len(this.Imei) == 0 || len(this.Mac) == 0 || len(this.Ip) == 0 {
		return &errors.InvalidParamsError
	}

	upsert := func(c *mgo.Collection) error {
		_, err := c.Upsert(bson.M{"imei": this.Imei}, this)
		return err
	}

	if err := withCollection(C_Router, upsert); err != nil {
		return &errors.DbError
	}

	return nil
}

func (this *Router) SetOnline(online bool) *errors.Error {
	if online && len(this.AccessToken) == 0 {
		return &errors.AccessError
	}

	change := bson.M{
		"$set": bson.M{
			"online":       online,
			"access_token": this.AccessToken,
		},
	}

	update := func(c *mgo.Collection) error {
		return c.Update(bson.M{"imei": this.Imei}, change)
	}

	if err := withCollection(C_Router, update); err != nil {
		return &errors.DbError
	}

	this.Online = online
	return nil
}
