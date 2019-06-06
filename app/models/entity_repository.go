package models

import (
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
func AddEntity(e Entity) (entity Entity, err error) {
	c := NewEntityCollection()
	defer c.Close()
	e.ID = bson.NewObjectId()
	e.CreatedAt = time.Now()
	return e, c.Session.Insert(e)
}

/**
GetEntities Get all Entities from database and returns
list of Entity on success
 */
func GetEntities(search string, categories []string) ([]Entity, error) {
	var (
		entities []Entity
		err      error
	)

	c := NewEntityCollection()
	defer c.Close()

	query := bson.M{
		"$text": bson.M{"$search": search},
	}
	if categories != nil && len(categories) != 0 {
		query["categories"] = bson.M{"$in": categories}
	}

	err = c.Session.Find(query).Select(bson.M{
		"score": bson.M{"$meta": "textScore"}}).Sort("$textScore:score").All(&entities)

	return entities, err
}

/**
GetEntity Get a Entity from database and returns
a Entity on success
 */
func GetEntity(id bson.ObjectId) (Entity, error) {
	var (
		entity Entity
		err    error
	)

	c := NewEntityCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"_id": id}).One(&entity)
	return entity, err
}

/**
GetEntity Get a Entity from database and returns
a Entity on success
 */
func GetEntityBy(attribute string, value string) (Entity, error) {
	var (
		entity Entity
		err    error
	)

	c := NewEntityCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{attribute: value}).One(&entity)
	return entity, err
}
