package models

import (
	"GIG/app/utilities/normalizers"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type NormalizedName struct {
	Id             bson.ObjectId `json:"id" bson:"_id"`
	SearchText     string        `json:"search_text" bson:"search_text"`
	NormalizedText string        `json:"normalized_text" bson:"normalized_text"`
	CreatedAt      time.Time     `json:"created_at" bson:"created_at"`
}

func (n NormalizedName) NewNormalizedName() NormalizedName {
	n.Id = bson.NewObjectId()
	n.CreatedAt = time.Now()
	return n
}

func (n NormalizedName) GetSearchText() string {
	return n.SearchText
}

func (n NormalizedName) SetSearchText(value string) NormalizedName {
	n.SearchText = normalizers.ProcessNameString(value)
	return n
}

func (n NormalizedName) GetNormalizedText() string {
	return n.NormalizedText
}

func (n NormalizedName) SetNormalizedText(value string) NormalizedName {
	n.NormalizedText = value
	return n
}

func (n NormalizedName) GetCreatedDate() time.Time {
	return n.CreatedAt
}
