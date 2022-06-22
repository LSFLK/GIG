package mongodb_official

import (
	"GIG/app/databases/index_manager"
	"context"
	"log"
	"os"
)

type MongoOfficialDatabaseHandler struct {
}

var service MongoOfficialDatabaseService
var Context = context.TODO()

func InitConnection(path string, dbName string, maxPool int) {
	if service.baseSession == nil {
		service.URL = path
		service.Database = dbName
		service.MaxPool = maxPool

		err := service.new()
		if err != nil {
			log.Println("error connecting to MongoDB database server:", service.URL)
			os.Exit(1)
		}
		index_manager.CreateDBIndexes(MongoOfficialIndexManager{})
	}
}

func DisconnectService() {
	err := service.client.Disconnect(Context)
	if err != nil {
		log.Fatal("error shutting down Mongo Client")
	}
}
