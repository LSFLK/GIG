package mongodb_official

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoOfficialDatabaseService struct {
	client  *mongo.Client
	path    string
	dbName  string
	maxPool int
}

func (s *MongoOfficialDatabaseService) new(path string, name string, maxPool int) error {
	var err error
	s.path, s.dbName, s.maxPool = path, name, maxPool
	log.Println("creating new mongodb client...")
	client, err := mongo.NewClient(options.Client().ApplyURI(service.path).SetMaxPoolSize(uint64(service.maxPool)))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(Context)
	if err != nil {
		return err
	}
	err = client.Ping(Context, nil)
	if err != nil {
		return err
	}
	log.Println("connected to mongodb successfully")
	service.client = client
	return nil
}
