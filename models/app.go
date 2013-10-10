// app
package models

import (
	"github.com/ginuerzh/mbbcloud/errors"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"time"
)

type App struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Version     string        `json:"version" bson:",omitempty"`
	Icon        string        `json:"icon" bson:",omitempty"`
	RUrl        string        `json:"url" bson:"rurl,omitempty"`
	CUrl        string        `json:"curl" bson:"-"`
	IUrl        string        `json:"iurl,omitempty" bson:"iurl,omitempty"`
	AUrl        string        `json:"aurl,omitempty" bson:"aurl,omitempty"`
	JPubTime    string        `json:"pub_time" bson:"-"`
	PubTime     time.Time     `json:"-" bson:"pub_time"`
	JUpdateTime string        `json:"update_time" bson:"-"`
	UpdateTime  time.Time     `json:"-" bson:"update_time"`
}

type AppList struct {
	apps []App
}

func (this *App) FindOneBy(key string, value interface{}) (e *errors.Error) {
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{key: value}).One(this)
	}

	if err := withCollection(C_App, query); err != nil {
		if err == mgo.ErrNotFound {
			e = &errors.FileNotFoundError
		} else {
			log.Println(err)
			e = &errors.DbError
		}
	}

	return
}

func (this *App) Save() *errors.Error {
	upsert := func(c *mgo.Collection) error {
		this.Id = bson.NewObjectId()
		return c.Insert(this)
	}

	if err := withCollection(C_App, upsert); err != nil {
		return &errors.DbError
	}

	return nil
}

func (this *AppList) find(selector interface{}, skip, limit int) *errors.Error {
	query := func(c *mgo.Collection) error {
		query := c.Find(selector).Skip(skip)
		if limit > 0 {
			query = query.Limit(limit)
		}
		return query.All(&this.apps)
	}
	if err := withCollection(C_App, query); err != nil {
		return &errors.DbError
	}
	return nil
}

func (this *AppList) FindAll(skip, limit int) *errors.Error {
	return this.find(nil, skip, limit)
}

func (this *AppList) Apps() []App {
	if this.apps == nil {
		return []App{}
	}
	return this.apps
}

func (this *AppList) Len() int {
	return len(this.apps)
}
