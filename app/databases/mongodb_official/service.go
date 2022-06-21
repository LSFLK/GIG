package mongodb_official

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Service struct {
	baseSession mongo.Session
	client      *mongo.Client
	queue       chan int
	URL         string
	Open        int
}

var service Service
var Context = context.TODO()

func (s *Service) New() error {
	var err error
	s.queue = make(chan int, MaxPool)
	for i := 0; i < MaxPool; i = i + 1 {
		s.queue <- 1
	}
	log.Println("creating new mongodb client...")
	client, err := mongo.NewClient(options.Client().ApplyURI(service.URL))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(Context)
	if err != nil {
		log.Fatal(err)
	}

	s.Open = 0
	s.client = client
	s.baseSession, err = client.StartSession()
	return err
}

func (s *Service) Session() *mongo.Session {
	<-s.queue
	s.Open++
	newSession := s.baseSession
	return &newSession
}

func (s *Service) Close(c *Collection) {
	s.queue <- 1
	s.Open--
}
