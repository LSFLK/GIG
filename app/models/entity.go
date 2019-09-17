package models

import (
	"GIG/commons"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Entity struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title      string        `json:"title" bson:"title"`
	ImageURL   string        `json:"image_url" bson:"image_url"`
	SourceURL  string        `json:"source_url" bson:"source_url"`
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
Compare if a given entity source is equal to this entity
 */
func (e Entity) SameSource(otherEntity Entity) bool {
	return e.SourceURL != "" && e.SourceURL == otherEntity.SourceURL
}

/**
Check if the entity has data
 */
func (e Entity) HasContent() bool {
	if len(e.Links) != 0 {
		return true
	}
	if len(e.Categories) != 0 {
		return true
	}
	if len(e.Attributes) != 0 {
		return true
	}
	return false
}

/**
Check if the entity has no title, data
 */
func (e Entity) IsNil() bool {
	if e.Title != "" {
		return false
	}
	return !e.HasContent()
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
Get an attribute
 */
func (e Entity) GetAttribute(attributeName string) (Attribute, error) {
	for _, attribute := range e.Attributes {
		if attribute.Name == attributeName {
			return attribute, nil
		}
	}
	return Attribute{}, errors.New("Attribute not found.")
}

/**
Add new link to entity
 */
func (e Entity) AddLink(entity Entity) Entity {
	if commons.StringInSlice(e.Links, entity.Title) {
		return e
	}
	if entity.Title != "" {
		e.Links = append(e.Links, entity.Title)
	}
	return e
}

/**
Add new category to entity
 */
func (e Entity) AddCategory(category string) Entity {
	if commons.StringInSlice(e.Categories, category) {
		return e
	}
	e.Categories = append(e.Categories, category)
	return e
}
