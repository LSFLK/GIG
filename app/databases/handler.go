package databases

import (
	"GIG/app/constants/mongo_drivers"
	"GIG/app/databases/mongodb"
	"GIG/app/databases/mongodb_official"
	"log"

	"github.com/revel/revel"
)

var driver, path, dbName string
var maxPool int

func LoadDatabaseHandler() {
	getDBConfig()
	var err error
	if err != nil {
		log.Fatal(err)
	}
	switch driver {
	case mongo_drivers.Mongodb:
		mongodb.LoadMongo()
	case mongo_drivers.MongodbOfficial:
		mongodb_official.InitConnection(path, dbName, maxPool)
	}
}

func CloseDatabaseHandler() {
	log.Println("shutting down database clients...")
	switch driver {
	case mongo_drivers.MongodbOfficial:
		mongodb_official.DisconnectService()
	}
}

func getDBConfig() {
	var found bool
	driver, found = revel.Config.String("mongo.driver")
	if !found {
		log.Fatal("MongoDB driver not configured")
	}

	maxPool = revel.Config.IntDefault("mongo.maxPool", 20)
	path, found = revel.Config.String("mongo.path")
	if !found {
		log.Fatal("MongoDB path not configured")
	}
	dbName, found = revel.Config.String("mongo.database")
	if !found {
		log.Fatal("MongoDB database not configured")
	}
}
