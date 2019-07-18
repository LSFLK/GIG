package mongodb

import (
	"github.com/revel/config"
	"github.com/revel/revel"
	"log"
)

func LoadMongo() {
	Config, err := config.LoadContext("app.conf",revel.CodePaths)
	if err != nil || Config == nil {
		log.Fatalf("%+v",err)
	}
	MaxPool = revel.Config.IntDefault("mongo.maxPool", 0)
	PATH,_ = revel.Config.String("mongo.path")
	DBNAME, _ = revel.Config.String("mongo.database")
	CheckAndInitServiceConnection()
}
