// model.go
package models

import (
	"github.com/astaxie/beego"
	"labix.org/v2/mgo"
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

	NSPrefix      = "mbbcloud:"
	NSRouters     = NSPrefix + "routers"
	NSRouter      = NSPrefix + "router:"
	NSRouterUsers = NSPrefix + "users:"
	NSRouterUser  = NSPrefix + "user:"
	NSMQ          = NSPrefix + "mq:"
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
