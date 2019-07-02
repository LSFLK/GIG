package repository

import (
	"GIG/app/models"
	"GIG/app/models/mongodb"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
)

func NewEntityCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("entities")
}

/**
AddEntity insert a new Entity into database and returns
last inserted entity on success.
 */
func AddEntity(entity models.Entity) (models.Entity, error) {
	entity.Title = strings.NewReplacer(
		"%", "",
		"/", "-",
		"~", "2",
	).Replace(entity.Title)
	entity.LoadedLinks = nil
	existingEntity, err := GetEntityBy("title", entity.Title)
	if entity.IsEqualTo(existingEntity) && existingEntity.HasContent() {
		return existingEntity, err
	}
	if !existingEntity.IsNil() && !existingEntity.HasContent() {
		entity.ID = existingEntity.ID
		entity.UpdatedAt = time.Now()
		entity.CreatedAt = existingEntity.CreatedAt
		err = UpdateEntity(entity)
	} else {
		fmt.Println("new")
		entity.ID = bson.NewObjectId()
		entity.UpdatedAt = time.Now()
		fmt.Println(entity)

		c := NewEntityCollection()
		defer c.Close()
		return entity, c.Session.Insert(entity)
	}

	return entity, err

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
	e.LoadedLinks = nil
	for _, link := range e.LinkIds {
		relatedEntity, relatedEntityErr := GetEntity(link)
		if relatedEntityErr == nil {
			e.LoadedLinks = append(e.LoadedLinks, relatedEntity)
		}
	}
	return e
}

/**
GetEntity Get a Entity from database and returns
a models. Entity on success
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
		"$set": e,
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
