package models

import (
	"GIG/app/utility"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Entity struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	Title      string        `json:"title" bson:"title"`
	Attributes []Attribute   `json:"attributes" bson:"attributes"`
	Links      []string      `json:"links" bson:"links"`
	Categories []string      `json:"categories" bson:"categories"`
	CreatedAt  time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at" bson:"updated_at"`
}

/**
Compare if a given entity is equal to this entity
 */
func (e Entity) IsEqualTo(otherEntity Entity) bool {
	return e.Title == otherEntity.Title
}

/**
Check if the entity has no data
 */
func (e Entity) IsEmpty() bool {
	if e.Title == "" {
		return false
	}
	if len(e.Links) != 0 {
		return false
	}
	if len(e.Categories) != 0 {
		return false
	}
	if len(e.Attributes) != 0 {
		return false
	}
	return true
}

/**
Add or update an existing attribute with a new value
 */
func (e Entity) SetAttribute(attributeName string, value Value) Entity {
	//iterate through all attributes
	var attributes []Attribute
	attributeFound := false
	for _, attribute := range e.Attributes {
		if attribute.Name == attributeName { //if attribute name matches an existing attribute
			attribute = attribute.SetValue(value) // append new value to the attribute
			attributeFound = true
		}
		attributes = append(attributes, attribute)
	}
	if !attributeFound { //else create new attribute and append value

		attribute := Attribute{Name: attributeName}.SetValue(value)
		attributes = append(attributes, attribute)
	}
	e.Attributes = attributes

	return e
}

/**
Add new link to entity
 */
func (e Entity) AddLink(link string) Entity {
	if utility.StringInSlice(e.Links, link) {
		return e
	}
	e.Links = append(e.Links, link)
	return e
}

/**
Add new category to entity
 */
func (e Entity) AddCategory(category string) Entity {
	if utility.StringInSlice(e.Categories, category) {
		return e
	}
	e.Categories = append(e.Categories, category)
	return e
}
