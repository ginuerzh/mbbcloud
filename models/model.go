// model.go
package models

import (
	"labix.org/v2/mgo"
)

var (
	DB *mgo.Database
)

const (
	C_Router = "routers"
	C_App    = "app"
	C_User   = "users"

	NSPrefix      = "mbbcloud:"
	NSRouters     = NSPrefix + "routers"
	NSRouter      = NSPrefix + "router:"
	NSRouterUsers = NSPrefix + "users:"
	NSRouterUser  = NSPrefix + "user:"
)
