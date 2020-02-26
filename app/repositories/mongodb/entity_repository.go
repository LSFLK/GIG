package mongodb

import (
	"GIG/app/databases/mongodb"
	"GIG/app/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct {
}

func NewEntityCollection() *mongodb.Collection {
	c := mongodb.NewCollectionSession("entities")
	textIndex := mgo.Index{
		Key: []string{"$text:title"},
		Weights: map[string]int{
			"title": 1,
		},
		Name: "textIndex",
	}
	titleIndex := mgo.Index{
		Key:    []string{"title"},
		Name:   "titleIndex",
		Unique: true,
	}
	c.Session.EnsureIndex(textIndex)
	c.Session.EnsureIndex(titleIndex)
	return c
}

/*
AddEntity insert a new Entity into database and returns
last inserted entity on success.
 */
func (r Repository) AddEntity(entity models.Entity) (models.Entity, error) {
	c := NewEntityCollection()
	defer c.Close()
	return entity, c.Session.Insert(entity)
}

/**
GetEntities Get all Entities where a given title is linked from
list of models.Entity on success
 */
func (r Repository) GetRelatedEntities(entity models.Entity, limit int) ([]models.Entity, error) {
	var (
		entities []models.Entity
		err      error
	)

	query := bson.M{}
	c := NewEntityCollection()
	defer c.Close()

	if entity.GetTitle() != "" {
		query["links"] = bson.M{"$in": []string{entity.GetTitle()}}

		// if the entity is not of primitive type
		if len(entities) == 0 {
			query["links"] = bson.M{"$in": []string{entity.GetTitle()}}
		}
	}
	err = c.Session.Find(query).Sort("-_id").Limit(limit).All(&entities)

	return entities, err
}

/**
GetEntities Get all Entities from database and returns
list of models.Entity on success
 */
func (r Repository) GetEntities(search string, categories []string, limit int) ([]models.Entity, error) {
	var (
		entities    []models.Entity
		err         error
		resultQuery *mgo.Query
	)

	query := bson.M{}
	c := NewEntityCollection()
	defer c.Close()

	if search != "" {
		query = bson.M{
			"$text": bson.M{"$search": search},
			//"attributes": bson.M{"$exists": true, "$not": bson.M{"$size": 0}},
		}
	}

	if categories != nil && len(categories) != 0 {
		query["categories"] = bson.M{"$all": categories}
	}

	// sort by search score for text indexed search, otherwise sort by latest first in category
	if search == "" {
		resultQuery = c.Session.Find(query).Sort("-updated_at")
	} else {
		resultQuery = c.Session.Find(query).Select(bson.M{
			"score": bson.M{"$meta": "textScore"}}).Sort("$textScore:score")
	}

	err = resultQuery.Limit(limit).All(&entities)

	return entities, err
}

/**
GetEntity Get a Entity from database and returns
a models. Entity on success
 */
func (r Repository) GetEntity(id bson.ObjectId) (models.Entity, error) {
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
func (r Repository) GetEntityBy(attribute string, value string) (models.Entity, error) {
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
func (r Repository) UpdateEntity(e models.Entity) error {
	c := NewEntityCollection()
	defer c.Close()

	err := c.Session.Update(bson.M{
		"_id": e.GetId(),
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

	err := c.Session.Remove(bson.M{"_id": e.GetId()})
	return err
}
