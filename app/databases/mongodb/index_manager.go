package mongodb

import (
	"GIG/app/databases/index_manager"
	"gopkg.in/mgo.v2"
	"sync"
)

type MongoLegacyIndexManager struct {
	index_manager.IndexManager
}

func (m MongoLegacyIndexManager) CreateEntityIndexes(wg *sync.WaitGroup) {
	c := NewCollectionSession("entities")
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

func (m MongoLegacyIndexManager) CreateNormalizeNameIndexes(wg *sync.WaitGroup) {
	c := NewCollectionSession("normalized_names")
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
	c := NewCollectionSession("users")
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
