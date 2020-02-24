package mongodb

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type valueModel struct {
	Type      string    `json:"type" bson:"type"`
	RawValue  string    `json:"raw_value" bson:"raw_value"`
	Source    string    `json:"source" bson:"source"`
	Date      time.Time `json:"date" bson:"date"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type dataModel struct {
	ID         bson.ObjectId                    `json:"id" bson:"_id,omitempty"`
	Title      string                           `json:"title" bson:"title"`
	ImageURL   string                           `json:"image_url" bson:"image_url"`
	SourceURL  string                           `json:"source_url" bson:"source_url"`
	Attributes map[string]map[string]valueModel `json:"attributes" bson:"attributes"`
	Links      []string                         `json:"links" bson:"links"`
	Categories []string                         `json:"categories" bson:"categories"`
	CreatedAt  time.Time                        `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time                        `json:"updated_at" bson:"updated_at"`
	Snippet    string                           `json:"snippet" bson:"snippet"`
}
