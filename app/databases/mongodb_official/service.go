package mongodb_official

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service struct {
	baseSession mongo.Session
	client      *mongo.Client
	queue       chan int
	URL         string
	Open        int
}

var service Service

func (s *Service) New() error {
	var err error
	s.queue = make(chan int, MaxPool)
	for i := 0; i < MaxPool; i = i + 1 {
		s.queue <- 1
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(s.URL))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

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
