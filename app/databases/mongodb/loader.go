package mongodb

import (
	"github.com/revel/revel"
)

func LoadMongo() {
	MaxPool = revel.Config.IntDefault("mongo.maxPool", 0)
	PATH,_ = revel.Config.String("mongo.path")
	DBNAME, _ = revel.Config.String("mongo.database")
	CheckAndInitServiceConnection()
}
