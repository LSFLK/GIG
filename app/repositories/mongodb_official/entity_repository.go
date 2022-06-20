package mongodb_official

import (
	"GIG/app/databases/mongodb_official"
	"GIG/app/repositories/constants"
	"github.com/lsflk/gig-sdk/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type EntityRepository struct {
}

func (e EntityRepository) newEntityCollection() *mongodb_official.Collection {
	c := mongodb_official.NewCollectionSession("entities")
	return c
}

/*
AddEntity insert a new Entity into database and returns
last inserted entity on success.
*/
func (e EntityRepository) AddEntity(entity models.Entity) (models.Entity, error) {
	c := e.newEntityCollection()
	defer c.Close()
	_, err := c.Collection.InsertOne(mongodb_official.Context, entity)
	return entity, err
}

func (e EntityRepository) GetEntityByPreviousTitle(title string, date time.Time) (models.Entity, error) {
	var (
		entity models.Entity
	)

	query := bson.M{
		"attributes.titles.values.value_string": title,
		"attributes.titles.values.date":         bson.M{"$lt": date.Add(time.Duration(1) * time.Second)},
	}

	c := e.newEntityCollection()
	defer c.Close()
	findOptions := options.FindOne()
	findOptions.SetSort(bson.D{{"attributes.titles.values.date", -1}})
	return entity, c.Collection.FindOne(mongodb_official.Context, query, findOptions).Decode(&entity)
}

/*
GetRelatedEntities - Get all Entities where a given title is linked from
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
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"updated_at", -1}}).
		SetLimit(int64(limit)).
		SetSkip(int64(offset))
	cursor, err := c.Collection.Find(mongodb_official.Context, query, findOptions)
	err = cursor.Decode(&entities)
	if err != nil {
		return entities, err
	}

	for _, item := range entities {
		log.Println(item.GetTitle())
	}
	return entities, err
}

/*
GetEntities - Get all Entities from database and returns
list of models.Entity on success
*/
func (e EntityRepository) GetEntities(search string, categories []string, limit int, offset int) ([]models.Entity, error) {
	var (
		entities []models.Entity
		err      error
		cursor   *mongo.Cursor
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
		query["categories"] = bson.M{"$all": categories}
	}

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit)).
		SetSkip(int64(offset))

	// sort by search score for text indexed search, otherwise sort by latest first in category
	if search == "" {
		findOptions.SetSort(bson.D{{"source_date", -1}})
		cursor, err = c.Collection.Find(mongodb_official.Context, query, findOptions)
	} else {
		findOptions.SetSort(bson.D{{"$textScore:score", 1}})
		cursor, err = c.Collection.Find(mongodb_official.Context, query, findOptions)
		//cursor.Select(bson.M{
		//	"score": bson.M{"$meta": "textScore"}}) TODO: check why select is used
	}
	err = cursor.Decode(&entities)
	log.Println(entities)
	return entities, err
}

/*
GetEntity - Get a Entity from database and returns
a models. Entity on success
*/
func (e EntityRepository) GetEntity(id string) (models.Entity, error) {
	var (
		entity models.Entity
		err    error
	)

	c := e.newEntityCollection()
	defer c.Close()

	cursor := c.Collection.FindOne(mongodb_official.Context, bson.M{"_id": id})
	cursor.Decode(&entity)
	return entity, err
}

/*
GetEntityBy -  Get a Entity from database by attribute value and returns
a models.Entity on success
*/
func (e EntityRepository) GetEntityBy(attribute string, value string) (models.Entity, error) {
	var (
		entity models.Entity
		err    error
	)

	c := e.newEntityCollection()
	defer c.Close()
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"updated_at", -1}})
	cursor := c.Collection.FindOne(mongodb_official.Context, bson.M{attribute: value})
	cursor.Decode(&entity)
	return entity, err
}

/*
UpdateEntity - update a Entity into database and returns
last nil on success.
*/
func (e EntityRepository) UpdateEntity(entity models.Entity) error {
	c := e.newEntityCollection()
	defer c.Close()
	filter := bson.D{{"_id", entity.GetId()}}
	update := bson.D{{"$set", entity}}
	_, err := c.Collection.UpdateOne(mongodb_official.Context, filter, update)
	return err
}

/*
DeleteEntity - Delete Entity from database and returns
last nil on success.
*/
func (e EntityRepository) DeleteEntity(entity models.Entity) error {
	c := e.newEntityCollection()
	defer c.Close()
	filter := bson.D{{"_id", entity.GetId()}}
	_, err := c.Collection.DeleteOne(mongodb_official.Context, filter)
	return err
}

/*
GetStats - Get entity states from the DB
*/
func (e EntityRepository) GetStats() (models.EntityStats, error) {
	var (
		entityStats models.EntityStats
		err         error
	)

	c := e.newEntityCollection()
	defer c.Close()

	// Get total number of entities
	entityCount, err := c.Collection.CountDocuments(mongodb_official.Context, nil)
	entityStats.EntityCount = int(entityCount)
	var linkCount []map[string]interface{}

	//Get category wise count
	categoryCountPipeline := []bson.M{
		{constants.UnwindAttribute: constants.CategoryAttribute},
		{constants.GroupAttribute: bson.M{
			"_id":            constants.CategoryAttribute,
			"category_count": bson.M{"$sum": 1}}},
		{constants.SortAttribute: bson.M{"category_count": -1}},
	}
	cursor, err := c.Collection.Aggregate(mongodb_official.Context, categoryCountPipeline)
	cursor.Decode(&entityStats.CategoryWiseCount)

	//Get category group wise count
	categoryGroupCountPipeline := []bson.M{
		{constants.UnwindAttribute: constants.CategoryAttribute},
		{constants.SortAttribute: bson.M{"categories": 1}},
		{constants.GroupAttribute: bson.M{"_id": "$_id", "sortedCategories": bson.M{"$push": constants.CategoryAttribute}}},
		{
			constants.GroupAttribute: bson.M{
				"_id":            "$sortedCategories",
				"category_count": bson.M{"$sum": 1}}},
		{constants.SortAttribute: bson.M{"category_count": -1}},
	}
	cursor, err = c.Collection.Aggregate(mongodb_official.Context, categoryGroupCountPipeline)
	cursor.Decode(&entityStats.CategoryGroupWiseCount)

	// Get total number of relations
	linkSumPipeline := []bson.M{{
		constants.GroupAttribute: bson.M{
			"_id":      "$link_sum",
			"link_sum": bson.M{"$sum": bson.M{"$size": "$links"}}}}}

	cursor, err = c.Collection.Aggregate(mongodb_official.Context, linkSumPipeline)
	cursor.Decode(&linkCount)
	entityStats.RelationCount, _ = linkCount[0]["link_sum"].(int)

	return entityStats, err
}
