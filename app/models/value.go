package models

import "time"

type Value struct {
	Type      string    `json:"type" bson:"type"` // type can be string, json, etc.
	RawValue  string    `json:"raw_value" bson:"raw_value"`
	StartDate time.Time `json:"start_date" bson:"start_date"`
	EndDate   time.Time `json:"end_date" bson:"end_date"`
}
