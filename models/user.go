// user
package models

import (
	"github.com/ginuerzh/mbbcloud/errors"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

type User struct {
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	Rank       int      `json:"rank"`
	LastAccess JsonTime `json:"last_access" bson:"last_access"`
}

func (this *User) FindOneBy(key string, value interface{}) (e *errors.Error) {
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{key: value}).One(this)
	}

	if err := withCollection(C_User, query); err != nil {
		if err == mgo.ErrNotFound {
			e = &errors.NotFoundError
		} else {
			e = &errors.DbError
		}
	}

	return
}

func (this *User) Save() *errors.Error {
	if len(this.Username) == 0 || len(this.Password) == 0 {
		return &errors.InvalidParamsError
	}

	this.LastAccess = NewJsonTime(bson.Now())
	upsert := func(c *mgo.Collection) error {
		_, err := c.Upsert(bson.M{"username": this.Username}, this)
		return err
	}

	if err := withCollection(C_User, upsert); err != nil {
		log.Println(err)
		return &errors.DbError
	}

	return nil
}
