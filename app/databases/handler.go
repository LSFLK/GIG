package databases

import (
	"GIG/app/constants/mongo_drivers"
	"GIG/app/databases/mongodb"
	"GIG/app/databases/mongodb_official"
	"log"

	"github.com/revel/revel"
)

func LoadDatabaseHandler() {
	driver, found := revel.Config.String("mongo.driver")
	if !found {
		log.Fatal("MongoDB driver not configured")
	}
	switch driver {
	case mongo_drivers.Mongodb:
		mongodb.LoadMongo()
	case mongo_drivers.MongodbOfficial:
		mongodb_official.LoadMongo()
	}

}
