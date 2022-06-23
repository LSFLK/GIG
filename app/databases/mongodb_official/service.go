package mongodb_official

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoOfficialDatabaseService struct {
	baseSession mongo.Session
	client      *mongo.Client
	queue       chan int
	open        int
	URL         string
	MaxPool     int
	Database    string
}

func (s MongoOfficialDatabaseService) new() error {
	var err error
	service.queue = make(chan int, service.MaxPool)
	for i := 0; i < service.MaxPool; i = i + 1 {
		service.queue <- 1
	}
	log.Println("creating new mongodb client...")
	client, err := mongo.NewClient(options.Client().ApplyURI(service.URL).SetMaxPoolSize(uint64(service.MaxPool)))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(Context)
	if err != nil {
		log.Fatal(err)
	}

	service.open = 0
	service.client = client
	service.baseSession, err = client.StartSession()
	return err
}

func (s MongoOfficialDatabaseService) Session() *mongo.Session {
	<-service.queue
	service.open++
	newSession := service.baseSession // create a copy of the base session
	return &newSession
}

func (s MongoOfficialDatabaseService) Close(c *Collection) {
	(*c.db.s).EndSession(Context)
	service.queue <- 1
	service.open--
}
