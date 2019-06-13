package models

import (
	"GIG/app/utility"
	"fmt"
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
Eager load entity related attributes
 */
func (e Entity) EagerLoad() Entity {
	/**
	iterate attributes and find objectIds and load entity Titles
	 */
	var attributes []Attribute
	for _, attribute := range e.Attributes {
		var values []Value
		for _, value := range attribute.Values {
			if value.Type == "objectId" {
				relatedEntity, relatedEntityErr := GetEntity(bson.ObjectIdHex(value.RawValue))
				fmt.Println(relatedEntity)
				if relatedEntityErr == nil {
					value.Type = "string"
					value.RawValue = relatedEntity.Title
				}
			}
			values = append(values, value)
		}
		attribute.Values = values
		attributes = append(attributes, attribute)
	}
	e.Attributes = attributes

	/**
	find Titles for Links
	 */

	//for _,Link := range entity.Links{
	//
	//}

	return e
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
