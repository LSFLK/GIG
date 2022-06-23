package mongodb_official

import (
	"GIG/app/databases/index_manager"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

var service MongoOfficialDatabaseService
var Context = context.TODO()

func InitConnection(path string, dbName string, maxPool int) {
	if service.client == nil {

		err := service.new(path, dbName, maxPool)
		if err != nil {
			log.Println("error connecting to MongoDB database server:", service.path)
			os.Exit(1)
		}
		index_manager.CreateDBIndexes(MongoOfficialIndexManager{})
	}
}

func GetCollection(name string) *mongo.Collection {
	return service.client.Database(service.dbName).Collection(name)
}

func DisconnectService() {
	err := service.client.Disconnect(Context)
	if err != nil {
		log.Fatal("error shutting down Mongo Client")
	}
}
