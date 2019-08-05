package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type NormalizedName struct {
	ID             bson.ObjectId `json:"id" bson:"_id"`
	SearchText     string        `json:"searchText" bson:"searchText"`
	NormalizedText string        `json:"normalizedText" bson:"normalizedText"`
	CreatedAt      time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at" bson:"updated_at"`
}