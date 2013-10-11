// app
package models

import (
	"github.com/ginuerzh/mbbcloud/errors"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
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
	PubTime     JsonTime      `json:"pub_time" bson:"pub_time"`
	UpdateTime  JsonTime      `json:"update_time" bson:"update_time"`
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

func (this *App) Update() *errors.Error {
	update := func(c *mgo.Collection) error {
		return c.UpdateId(this.Id, this)
	}

	if err := withCollection(C_App, update); err != nil {
		return &errors.DbError
	}

	return nil
}

func (this *App) Delete() *errors.Error {
	remove := func(c *mgo.Collection) error {
		if err := this.FindOneBy("_id", this.Id); err != nil {
			return nil
		}
		file := File{}
		if len(this.Icon) > 0 {
			file.Fid = this.Icon
			file.Delete()
		}
		if len(this.RUrl) > 0 {
			file.Fid = this.RUrl
			file.Delete()
		}
		if len(this.IUrl) > 0 {
			file.Fid = this.IUrl
			file.Delete()
		}
		if len(this.AUrl) > 0 {
			file.Fid = this.AUrl
			file.Delete()
		}
		c.RemoveId(this.Id)
		return nil
	}

	if err := withCollection(C_App, remove); err != nil {
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
