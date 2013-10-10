// model.go
package models

import (
	"github.com/astaxie/beego"
	"labix.org/v2/mgo"
	//"log"
	"strconv"
	"time"
)

var (
	session *mgo.Session
	db      = beego.AppConfig.String("mongodb")
)

const (
	C_Router = "routers"
	C_App    = "apps"
	C_User   = "users"
	C_File   = "files"

	JsonTimeFormat = "2006-01-02 15:04:05"
)

// these codes are inspired by Denis Papathanasiou's post:
// http://denis.papathanasiou.org/2012/10/14/go-golang-and-mongodb-using-mgo/
func getSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.Dial(beego.AppConfig.String("mongourl"))
		if err != nil {
			panic(err) // no, not really
		}
	}
	return session.Clone()
}

func withCollection(collection string, f func(*mgo.Collection) error) error {
	s := getSession()
	defer s.Close()
	c := s.DB(db).C(collection)
	return f(c)
}

type JsonTime struct {
	time.Time
}

func NewJsonTime(t time.Time) JsonTime {
	return JsonTime{t}
}

func (t JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(t.Format(JsonTimeFormat))), nil
}

func (t *JsonTime) UnmarshalJSON(s []byte) (err error) {
	q, err := strconv.Unquote(string(s))
	if err != nil {
		return err
	}

	t.Time, err = time.Parse(JsonTimeFormat, q)
	return
}
