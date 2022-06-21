package mongodb_official

import (
	"GIG/app/databases/index_manager"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
)

type MongoOfficialIndexManager struct {
	index_manager.IndexManager
}

func (m MongoOfficialIndexManager) CreateEntityIndexes(wg *sync.WaitGroup) {
	c := NewCollectionSession("entities")
	textIndex := mongo.IndexModel{
		Keys:    bson.D{{"title", "text"}, {"search_text", "text"}},
		Options: options.Index().SetName("textIndex"),
	}
	titleIndex := mongo.IndexModel{
		Keys:    bson.D{{"title", 1}},
		Options: options.Index().SetName("titleIndex").SetUnique(true),
	}
	_, err := c.Collection.Indexes().CreateMany(Context, []mongo.IndexModel{textIndex, titleIndex})
	if err != nil {
		log.Fatal("error creating entity indexes:", err)
	}
	wg.Done()
}

func (m MongoOfficialIndexManager) CreateNormalizeNameIndexes(wg *sync.WaitGroup) {
	c := NewCollectionSession("normalized_names")
	textIndex := mongo.IndexModel{
		Keys:    bson.D{{"search_text", "text"}},
		Options: options.Index().SetName("textIndex").SetUnique(true),
	}
	searchTextIndex := mongo.IndexModel{
		Keys:    bson.D{{"search_text", 1}},
		Options: options.Index().SetName("searchTextIndex").SetUnique(true),
	}
	_, err := c.Collection.Indexes().CreateMany(Context, []mongo.IndexModel{textIndex, searchTextIndex})
	if err != nil {
		log.Fatal("error creating normalization indexes:", err)
	}
	wg.Done()
}
func (m MongoOfficialIndexManager) CreateUserIndexes(wg *sync.WaitGroup) {
	c := NewCollectionSession("users")
	userIndex := mongo.IndexModel{
		Keys:    bson.D{{"name", 1}},
		Options: options.Index().SetName("userIndex").SetUnique(true),
	}
	emailIndex := mongo.IndexModel{
		Keys:    bson.D{{"email", 1}},
		Options: options.Index().SetName("emailIndex").SetUnique(true),
	}
	_, err := c.Collection.Indexes().CreateMany(Context, []mongo.IndexModel{userIndex, emailIndex})
	if err != nil {
		log.Fatal("error creating user indexes:", err)
	}
	wg.Done()
}
