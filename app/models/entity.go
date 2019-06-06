package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Entity struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	SourceID   string        `json:"sourceId" bson:"sourceId"`
	Title      string        `json:"title" bson:"title"`
	Attributes []Attribute   `json:"attributes" bson:"attributes"`
	Links      []string      `json:"links" bson:"links"`
	Categories []string      `json:"categories" bson:"categories"`
	CreatedAt  time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at" bson:"updated_at"`
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
	e.Links = append(e.Links, link)
	return e
}

/**
Add new category to entity
 */
func (e Entity) AddCategory(link string) Entity {
	e.Categories = append(e.Categories, link)
	return e
}

/**
UpdateEntity update a Entity into database and returns
last nil on success.
 */
func (e Entity) UpdateEntity() error {
	c := NewEntityCollection()
	defer c.Close()

	err := c.Session.Update(bson.M{
		"_id": e.ID,
	}, bson.M{
		"$set": bson.M{
			"title": e.Title, "updatedAt": time.Now()},
	})
	return err
}

/**
DeleteEntity Delete Entity from database and returns
last nil on success.
 */
func (e Entity) DeleteEntity() error {
	c := NewEntityCollection()
	defer c.Close()

	err := c.Session.Remove(bson.M{"_id": e.ID})
	return err
}
