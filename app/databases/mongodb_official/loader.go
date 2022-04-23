package mongodb_official

import (
	"github.com/revel/revel"
)

var MaxPool int
var PATH string
var DBNAME string

func LoadMongo() {
	MaxPool = revel.Config.IntDefault("mongo.maxPool", 0)
	PATH, _ = revel.Config.String("mongo.path")
	DBNAME, _ = revel.Config.String("mongo.database")
	CheckAndInitServiceConnection()
}
