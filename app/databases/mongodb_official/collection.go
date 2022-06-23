package mongodb_official

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Collection struct {
	db         *Database
	name       string
	Collection *mongo.Collection
}

func (c *Collection) Connect() {
	collection := *c.db.database.Collection(c.name)
	c.Collection = &collection
}

func (c *Collection) GetSession() *mongo.Session {
	return c.db.s
}

func NewCollectionSession(name string) *Collection {
	var c = Collection{
		db:   newDBSession(service.Database),
		name: name,
	}
	c.Connect()
	return &c
}

func (c *Collection) Close() {
	service.Close(c)
}
