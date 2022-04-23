package databases

import (
	"GIG/app/databases/mongodb"
	"GIG/app/databases/mongodb_official"
	"log"
	"os"

	"github.com/revel/revel"
)

const (
	Mongodb         = "mongodb"
	MongodbOfficial = "mongodb-official"
)

func LoadDatabaseHandler() {
	driver, err := revel.Config.String("mongo.driver")
	if err {
		log.Println("MongoDB driver not configured", err)
		os.Exit(1)
	}
	switch driver {
	case Mongodb:
		mongodb.LoadMongo()
	case MongodbOfficial:
		mongodb_official.LoadMongo()
	}

}
