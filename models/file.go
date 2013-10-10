// file
package models

import (
	"github.com/ginuerzh/mbbcloud/errors"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type File struct {
	Fid         string
	Name        string `bson:"filename"`
	Size        int    `bson:"length"`
	Md5         string
	ContentType string    `bson:"contentType"`
	UploadDate  time.Time `bson:"uploadDate"`
}

func (this *File) FindOneBy(key string, value interface{}) (e *errors.Error) {
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{key: value}).One(this)
	}

	if err := withCollection(C_File, query); err != nil {
		if err == mgo.ErrNotFound {
			e = &errors.UserNotFoundError
		} else {
			e = &errors.DbError
		}
	}

	return
}

func (this *File) Save() *errors.Error {
	upsert := func(c *mgo.Collection) error {
		_, err := c.Upsert(bson.M{"fid": this.Fid}, this)
		return err
	}

	if err := withCollection(C_File, upsert); err != nil {
		return &errors.DbError
	}

	return nil
}

func (this *File) Exists() bool {
	count := 0
	query := func(c *mgo.Collection) error {
		var err error
		count, err = c.Find(bson.M{"fid": this.Fid}).Count()
		return err
	}

	if withCollection(C_File, query); count > 0 {
		return true
	}

	return false
}

func (this *File) Delete() *errors.Error {
	remove := func(c *mgo.Collection) error {
		return c.Remove(bson.M{"fid": this.Fid})
	}

	if err := withCollection(C_File, remove); err != nil {
		if err != mgo.ErrNotFound {
			return &errors.DbError
		}
	}
	return nil
}
