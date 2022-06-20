package mongodb_official

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func CreateDBIndexes() {
	c := NewCollectionSession("entities")
	textIndex := mongo.IndexModel{
		Keys:    bson.D{{"$text:title", 1}, {"$text:search_text", 1}},
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
}
