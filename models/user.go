// user
package models

import (
	"github.com/ginuerzh/mbbcloud/errors"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type User struct {
	Imei      string `json:"id"`
	Mac       string
	Ip        string
	Download  uint64
	Upload    uint64
	LoginTime time.Time `bson:"login_time"`
}

func (this *User) Load() (e *errors.Error) {
	c := DB.C(C_User)

	if err := c.Find(bson.M{"imei": this.Imei}).One(this); err != nil {
		if err == mgo.ErrNotFound {
			e = &errors.UserNotFoundError
		} else {
			e = &errors.DbError
		}
	}
	return
}

func (this *User) Exists() (b bool) {
	c := DB.C(C_User)

	count, _ := c.Find(bson.M{"imei": this.Imei}).Count()

	if count > 0 {
		b = true
	}

	return
}

func (this *User) Save() (e *errors.Error) {
	c := DB.C(C_User)

	if len(this.Imei) == 0 || len(this.Mac) == 0 || len(this.Ip) == 0 {
		return &errors.InvalidParamsError
	}

	this.LoginTime = bson.Now()

	if _, err := c.Upsert(bson.M{"imei": this.Imei}, this); err != nil {
		e = &errors.DbError
	}

	return
}
