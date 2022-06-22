package mongodb

import (
	"GIG/app/constants/database"
	"GIG/app/databases/interfaces"
	"gopkg.in/mgo.v2"
	"sync"
)

type MongoLegacyIndexManager struct {
	interfaces.IndexManagerInterface
}

func (m MongoLegacyIndexManager) CreateEntityIndexes(wg *sync.WaitGroup) {
	c := NewCollectionSession(database.EntityCollection)
	textIndex := mgo.Index{
		Key: []string{"$text:title", "$text:search_text"},
		Weights: map[string]int{
			"title":       1,
			"search_text": 1,
		},
		Name: "textIndex",
	}
	titleIndex := mgo.Index{
		Key:    []string{"title"},
		Name:   "titleIndex",
		Unique: true,
	}
	c.Collection.EnsureIndex(textIndex)
	c.Collection.EnsureIndex(titleIndex)
	wg.Done()
}

func (m MongoLegacyIndexManager) CreateNormalizedNameIndexes(wg *sync.WaitGroup) {
	c := NewCollectionSession(database.NormalizedNameCollection)
	textIndex := mgo.Index{
		Key: []string{"$text:search_text"},
		Weights: map[string]int{
			"search_text": 1,
		},
		Name:   "textIndex",
		Unique: true,
	}
	searchTextIndex := mgo.Index{
		Key:    []string{"search_text"},
		Name:   "searchTextIndex",
		Unique: true,
	}
	c.Collection.EnsureIndex(textIndex)
	c.Collection.EnsureIndex(searchTextIndex)
	wg.Done()
}
func (m MongoLegacyIndexManager) CreateUserIndexes(wg *sync.WaitGroup) {
	c := NewCollectionSession(database.UserCollection)
	userIndex := mgo.Index{
		Key:    []string{"name"},
		Name:   "userIndex",
		Unique: true,
	}
	emailIndex := mgo.Index{
		Key:    []string{"email"},
		Name:   "emailIndex",
		Unique: true,
	}
	c.Collection.EnsureIndex(userIndex)
	c.Collection.EnsureIndex(emailIndex)
	wg.Done()
}
