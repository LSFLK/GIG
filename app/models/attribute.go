package models

type Attribute struct {
	Name     string  `json:"name" bson:"name"`
	ValueObj []Value `json:"raw_value" bson:"raw_value"`
}
