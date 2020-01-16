package mongodb

import (
	"GIG/app/databases/mongodb"
	"GIG/app/models"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

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
func AddEntity(entity models.Entity) (models.Entity, error) {
	existingEntity, err := GetEntityBy("title", entity.Title)

	if err != nil {
		return models.Entity{}, err
	}

	if entity.UpdatedAt.IsZero() {
		entity.UpdatedAt = time.Now()
	}
	entity = entity.SetSnippet()

	//if an entity exists
	if entity.IsEqualTo(existingEntity) {
		//if the entity has a "new_title" attribute use it to change the entity title
		newTitleAttribute, err := entity.GetAttribute("new_title")

		if err == nil {
			fmt.Println("entity title modification found.", existingEntity.GetTitle(), "->", newTitleAttribute.GetValue().RawValue)
			existingEntity = existingEntity.SetTitle(newTitleAttribute.GetValue())
		}

		// merge links
		existingEntity = existingEntity.AddLinks(entity.Links)
		// merge categories
		existingEntity = existingEntity.AddCategories(entity.Categories)
		// merge attributes
		for _, attribute := range entity.Attributes {
			if attribute.Name != "new_title" {
				entityAttribute, _ := entity.GetAttribute(attribute.Name)
				existingEntity = existingEntity.SetAttribute(attribute.Name, entityAttribute.GetValue())
			}
		}
		// set updated date
		existingEntity.UpdatedAt = time.Now()
		fmt.Println("entity exists. updated", entity.Title)

		return existingEntity, UpdateEntity(existingEntity)
	} else {
		// if no entity exist
		entity.ID = bson.NewObjectId()
		entity.CreatedAt = time.Now()
		c := NewEntityCollection()
		defer c.Close()
		fmt.Println("creating new entity", entity.Title)
		return entity, c.Session.Insert(entity)
	}

}

/**
GetEntities Get all Entities where a given title is linked from
list of models.Entity on success
 */
func GetRelatedEntities(entity models.Entity, limit int) ([]models.Entity, error) {
	var (
		entities []models.Entity
		err      error
	)

	query := bson.M{}
	c := NewEntityCollection()
	defer c.Close()

	if entity.Title != "" {
		query["links"] = bson.M{"$in": []string{entity.Title}}

		// if the entity is not of primitive type
		if len(entities) == 0 {
			query["links"] = bson.M{"$in": entity.Links}
		}
	}
	err = c.Session.Find(query).Sort("-_id").Limit(limit).All(&entities)

	return entities, err
}

/**
GetEntities Get all Entities from database and returns
list of models.Entity on success
 */
func GetEntities(search string, categories []string, limit int) ([]models.Entity, error) {
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
