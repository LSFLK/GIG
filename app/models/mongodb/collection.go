package mongodb

import (
	"gopkg.in/mgo.v2"
)

type Collection struct {
	db      *Database
	name    string
	Session *mgo.Collection
}

func (c *Collection) Connect() {
	session := *c.db.session.C(c.name)
	index := mgo.Index{
		Key: []string{"$text:title", "$text:content", "$text:categories"},
		Weights: map[string]int{
			"title":   2,
			"content": 9,
			"categories": 1,
		},
		Name: "textIndex",
	}
	session.EnsureIndex(index)
	c.Session = &session
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

