package databases

import (
	"GIG/app/databases/mongodb"
	"GIG/app/databases/mongodb_official"
	"log"

	"github.com/revel/revel"
)

const (
	Mongodb         = "mongodb"
	MongodbOfficial = "mongodb-official"
)

func LoadDatabaseHandler() {
	driver, found := revel.Config.String("mongo.driver")
	log.Println(found)
	if !found {
		log.Fatal("MongoDB driver not configured")
	}
	switch driver {
	case Mongodb:
		mongodb.LoadMongo()
	case MongodbOfficial:
		mongodb_official.LoadMongo()
	}

}
