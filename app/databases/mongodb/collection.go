package mongodb

import (
	"gopkg.in/mgo.v2"
)

type Collection struct {
	db         *Database
	name       string
	Collection *mgo.Collection
}

func (c *Collection) Connect() {
	collection := *c.db.database.C(c.name)
	c.Collection = &collection
}

func NewCollectionSession(name string) *Collection {
	var c = Collection{
		db:   newDBSession(DBNAME),
		name: name,
	}
	c.Connect()
	return &c
}

func (c *Collection) Close() {
	service.Close(c)
}
