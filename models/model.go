// model.go
package models

import (
	"github.com/garyburd/redigo/redis"
	"labix.org/v2/mgo"
)

var (
	DB    *mgo.Database
	Redis redis.Conn
)

const (
	C_User = "User"

	RedisNSUserOnline = "user:online"
)
