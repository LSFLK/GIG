package repository

import (
	"GIG/app/models"
	"GIG/app/models/mongodb"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func NewEntityCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("entities")
}

/**
AddEntity insert a new Entity into database and returns
last inserted entity on success.
 */
func AddEntity(e models.Entity) (entity models.Entity, err error) {
	c := NewEntityCollection()
	defer c.Close()
	e.ID = bson.NewObjectId()
	e.CreatedAt = time.Now()
	return e, c.Session.Insert(e)
}

/**
GetEntities Get all Entities from database and returns
list of models.Entity on success
 */
func GetEntities(search string, categories []string) ([]models.Entity, error) {
	var (
		entities []models.Entity
		err      error
	)

	query := bson.M{}
	c := NewEntityCollection()
	defer c.Close()

	if search != "" {
		query = bson.M{
			"$text": bson.M{"$search": search},
		}
	}

	if categories != nil && len(categories) != 0 {
		query["categories"] = bson.M{"$all": categories}
	}

	err = c.Session.Find(query).Select(bson.M{
		"score": bson.M{"$meta": "textScore"}}).Sort("$textScore:score").Limit(10).All(&entities)

	return entities, err
}

/**
Eager load entity related attributes
 */
func EagerLoad(e models.Entity) models.Entity {
	/**
	iterate attributes and find objectIds and load entity Titles
	 */
	var attributes []models.Attribute
	for _, attribute := range e.Attributes {
		var values []models.Value
		for _, value := range attribute.Values {
			if value.Type == "objectId" {
				relatedEntity, relatedEntityErr := GetEntity(bson.ObjectIdHex(value.RawValue))
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
	var links []string
	for _, link := range e.Links {
		relatedEntity, relatedEntityErr := GetEntity(bson.ObjectIdHex(link))
		if relatedEntityErr == nil {
			links = append(links, relatedEntity.Title)
		}
	}
	e.Links = links

	return e
}

/**
GetEntity Get a Entity from database and returns
a models.Entity on success
 */
func GetEntity(id bson.ObjectId) (models.Entity, error) {
	var (
		entity models.Entity
		err    error
	)

	c := NewEntityCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"_id": id}).One(&entity)
	return entity, err
}

/**
GetEntity Get a Entity from database and returns
a models.Entity on success
 */
func GetEntityBy(attribute string, value string) (models.Entity, error) {
	var (
		entity models.Entity
		err    error
	)

	c := NewEntityCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{attribute: value}).One(&entity)
	return entity, err
}

/**
UpdateEntity update a Entity into database and returns
last nil on success.
 */
func UpdateEntity(e models.Entity) error {
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
func DeleteEntity(e models.Entity) error {
	c := NewEntityCollection()
	defer c.Close()

	err := c.Session.Remove(bson.M{"_id": e.ID})
	return err
}
