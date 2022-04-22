package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Service struct {
	baseSession *mongo.Session
	queue       chan int
	URL         string
	Open        int
	client      *mongo.Client
}

var service Service

func (s *Service) New() error {
	var err error
	s.queue = make(chan int, MaxPool)
	for i := 0; i < MaxPool; i = i + 1 {
		s.queue <- 1
	}
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(s.URL).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	defer s.client.Disconnect(ctx)

	err = s.client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Session() *mongo.Session {
	<-s.queue
	s.Open++
	return s.baseSession
}

func (s *Service) Close(c *Collection) {
	c.db.s.Close()
	s.queue <- 1
	s.Open--
}
