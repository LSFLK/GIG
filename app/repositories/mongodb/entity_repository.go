package mongodb

import (
	"GIG/app/databases/mongodb"
	"GIG-SDK//models"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type EntityRepository struct {
}

func (e EntityRepository) newEntityCollection() *mongodb.Collection {
	c := mongodb.NewCollectionSession("entities")
	textIndex := mgo.Index{
		Key: []string{"$text:title"},
		Weights: map[string]int{
			"title": 1,
		},
		Name:   "textIndex",
		Unique: true,
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
func (e EntityRepository) AddEntity(entity models.Entity) (models.Entity, error) {
	c := e.newEntityCollection()
	defer c.Close()
	return entity, c.Session.Insert(entity)
}

func (e EntityRepository) GetEntityByPreviousState(title string, date time.Time) ([]models.Entity, error) {
	var (
		entities []models.Entity
		err      error
	)

	query := bson.M{"attributes.titles.values.value_string": title}

	c := e.newEntityCollection()
	defer c.Close()

	err = c.Session.Find(query).Sort("-updated_at").All(&entities)
	return entities, err
}

/**
GetEntities Get all Entities where a given title is linked from
list of models.Entity on success
 */
func (e EntityRepository) GetRelatedEntities(entity models.Entity, limit int, offset int) ([]models.Entity, error) {
	var (
		entities []models.Entity
		err      error
	)

	query := bson.M{}
	c := e.newEntityCollection()
	defer c.Close()

	entityTitle := entity.GetTitle()
	if entityTitle != "" {
		query = bson.M{"links.title": bson.M{"$in": append(entity.GetLinkTitles(), entity.GetTitle())}}
	}
	fmt.Println(query)
	err = c.Session.Find(query).Sort("-updated_at").Skip(offset).Limit(limit).All(&entities)

	for _, item := range entities {
		fmt.Println(item.GetTitle())
	}
	return entities, err
}

/**
GetEntities Get all Entities from database and returns
list of models.Entity on success
 */
func (e EntityRepository) GetEntities(search string, categories []string, limit int, offset int) ([]models.Entity, error) {
	var (
		entities    []models.Entity
		err         error
		resultQuery *mgo.Query
	)

	query := bson.M{}
	c := e.newEntityCollection()
	defer c.Close()

	if search != "" {
		query = bson.M{
			"$text": bson.M{"$search": search},
			//"attributes": bson.M{"$exists": true, "$not": bson.M{"$size": 0}},
		}
	}

	if categories != nil && len(categories) != 0 {
		query["categories"] = bson.M{"$in": categories}
		//query["categories"] = bson.M{"$all": categories}
	}

	// sort by search score for text indexed search, otherwise sort by latest first in category
	if search == "" {
		resultQuery = c.Session.Find(query).Sort("-updated_at")
	} else {
		resultQuery = c.Session.Find(query).Select(bson.M{
			"score": bson.M{"$meta": "textScore"}}).Sort("$textScore:score")
	}

	err = resultQuery.Skip(offset).Limit(limit).All(&entities)

	return entities, err
}

/**
GetEntity Get a Entity from database and returns
a models. Entity on success
 */
func (e EntityRepository) GetEntity(id bson.ObjectId) (models.Entity, error) {
	var (
		entity models.Entity
		err    error
	)

	c := e.newEntityCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"_id": id}).One(&entity)
	return entity, err
}

/**
GetEntity Get a Entity from database and returns
a models.Entity on success
 */
func (e EntityRepository) GetEntityBy(attribute string, value string) (models.Entity, error) {
	var (
		entity models.Entity
		err    error
	)

	c := e.newEntityCollection()
	defer c.Close()
	err = c.Session.Find(bson.M{attribute: value}).Sort("-updated_at").One(&entity)
	return entity, err
}

/**
UpdateEntity update a Entity into database and returns
last nil on success.
 */
func (e EntityRepository) UpdateEntity(entity models.Entity) error {
	c := e.newEntityCollection()
	defer c.Close()

	err := c.Session.Update(bson.M{
		"_id": entity.GetId(),
	}, bson.M{
		"$set": entity,
	})
	return err
}

/**
DeleteEntity Delete Entity from database and returns
last nil on success.
 */
func (e EntityRepository) DeleteEntity(entity models.Entity) error {
	c := e.newEntityCollection()
	defer c.Close()

	err := c.Session.Remove(bson.M{"_id": entity.GetId()})
	return err
}
