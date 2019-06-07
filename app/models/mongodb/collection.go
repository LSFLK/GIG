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
	textIndex := mgo.Index{
		Key: []string{"$text:title"},
		Weights: map[string]int{
			"title":   1,
		},
		Name: "textIndex",
	}
	titleIndex := mgo.Index{
		Key: []string{"title:1"},
		Name: "titleIndex",
	}
	session.EnsureIndex(textIndex)
	session.EnsureIndex(titleIndex)
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

