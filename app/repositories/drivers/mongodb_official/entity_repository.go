package mongodb_official

import (
	"GIG/app/constants/database"
	"GIG/app/databases/mongodb_official"
	models2 "GIG/app/models"
	"GIG/app/repositories/constants"
	"GIG/app/repositories/interfaces"
	"github.com/lsflk/gig-sdk/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type EntityRepository struct {
	interfaces.EntityRepositoryInterface
}

func (e EntityRepository) newEntityCollection() *mongo.Collection {
	return mongodb_official.GetCollection(database.EntityCollection)
}

/*
AddEntity insert a new Entity into database and returns
last inserted entity on success.
*/
func (e EntityRepository) AddEntity(entity models.Entity) (models.Entity, error) {
	var err error
	collection := e.newEntityCollection()
	session, err := mongodb_official.GetSession()
	defer (*session).EndSession(mongodb_official.Context)
	if err = (*session).StartTransaction(); err != nil {
		return entity, err
	}
	if err = mongo.WithSession(mongodb_official.Context, *session, func(sc mongo.SessionContext) error {
		_, err := collection.InsertOne(mongodb_official.Context, entity)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return entity, err
	}
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
	findOptions := options.FindOne().
		SetSort(bson.D{{"attributes.titles.values.date", -1}})
	return entity, c.FindOne(mongodb_official.Context, query, findOptions).Decode(&entity)
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

	entityTitle := entity.GetTitle()
	if entityTitle != "" {
		query = bson.M{"links.title": bson.M{"$in": append(entity.GetLinkTitles(), entity.GetTitle())}}
	}
	findOptions := options.Find().
		SetSort(bson.D{{"updated_at", -1}}).
		SetLimit(int64(limit)).
		SetSkip(int64(offset))
	cursor, err := c.Find(mongodb_official.Context, query, findOptions)
	if err != nil {
		return entities, err
	}
	err = cursor.All(mongodb_official.Context, &entities)
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

	query := bson.D{}
	c := e.newEntityCollection()

	if search != "" {
		query = bson.D{
			{"$text", bson.D{{"$search", search}}},
		}
	}

	if categories != nil && len(categories) != 0 {
		query = bson.D{
			{"categories", bson.D{{"$all", categories}}},
		}
	}

	findOptions := options.Find().
		SetLimit(int64(limit)).
		SetSkip(int64(offset))

	// sort by search score for text indexed search, otherwise sort by latest first in category
	if search == "" {
		findOptions.SetSort(bson.D{{"source_date", -1}})
		cursor, err = c.Find(mongodb_official.Context, query, findOptions)
	} else {
		cursor, err = c.Find(mongodb_official.Context, query, findOptions)
	}
	if err != nil {
		return entities, err
	}
	if err = cursor.All(mongodb_official.Context, &entities); err != nil {
		return entities, err
	}
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

	cursor := c.FindOne(mongodb_official.Context, bson.M{"_id": id})
	err = cursor.Decode(&entity)
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
	findOptions := options.FindOne().
		SetSort(bson.D{{"updated_at", -1}})
	cursor := c.FindOne(mongodb_official.Context, bson.M{attribute: value}, findOptions)
	err = cursor.Decode(&entity)
	return entity, err
}

/*
UpdateEntity - update a Entity into database and returns
last nil on success.
*/
func (e EntityRepository) UpdateEntity(entity models.Entity) error {
	filter := bson.D{{"_id", entity.GetId()}}
	update := bson.D{{"$set", entity}}
	collection := e.newEntityCollection()
	session, err := mongodb_official.GetSession()
	defer (*session).EndSession(mongodb_official.Context)
	if err = (*session).StartTransaction(); err != nil {
		return err
	}
	if err = mongo.WithSession(mongodb_official.Context, *session, func(sc mongo.SessionContext) error {
		_, err := collection.UpdateOne(mongodb_official.Context, filter, update)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return err
}

/*
DeleteEntity - Delete Entity from database and returns
last nil on success.
*/
func (e EntityRepository) DeleteEntity(entity models.Entity) error {
	c := e.newEntityCollection()
	filter := bson.D{{"_id", entity.GetId()}}
	_, err := c.DeleteOne(mongodb_official.Context, filter)
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
	entityStats.CreatedAt = time.Now()

	c := e.newEntityCollection()

	// Get total number of entities
	entityCount, err := c.CountDocuments(mongodb_official.Context, bson.M{})
	entityStats.EntityCount = int(entityCount)
	var linkCount []map[string]int32

	//Get category wise count
	categoryCountPipeline := []bson.M{
		{constants.UnwindAttribute: constants.CategoryAttribute},
		{constants.GroupAttribute: bson.M{
			"_id":            constants.CategoryAttribute,
			"category_count": bson.M{"$sum": 1}}},
		{constants.SortAttribute: bson.M{"category_count": -1}},
	}
	cursor, err := c.Aggregate(mongodb_official.Context, categoryCountPipeline)
	if err != nil {
		return entityStats, err
	}
	err = cursor.All(mongodb_official.Context, &entityStats.CategoryWiseCount)
	if err != nil {
		return entityStats, err
	}

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
	cursor, err = c.Aggregate(mongodb_official.Context, categoryGroupCountPipeline)
	if err != nil {
		return entityStats, err
	}
	err = cursor.All(mongodb_official.Context, &entityStats.CategoryGroupWiseCount)
	if err != nil {
		return entityStats, err
	}

	// Get total number of relations
	linkSumPipeline := []bson.M{
		{"$match": bson.M{"links": bson.M{"$ne": nil}}},
		{constants.GroupAttribute: bson.M{
			"_id":      "$link_sum",
			"link_sum": bson.M{"$sum": bson.M{"$size": "$links"}}}}}

	cursor, err = c.Aggregate(mongodb_official.Context, linkSumPipeline)
	if err != nil {
		return entityStats, err
	}
	err = cursor.All(mongodb_official.Context, &linkCount)
	if err != nil {
		return entityStats, err
	}
	entityStats.RelationCount = int(linkCount[0]["link_sum"])

	return entityStats, err
}

/*
GetGraph - Get the entity relations summary for graph visualization
*/
func (e EntityRepository) GetGraph() (graph map[string]models2.GraphArray, err error) {
	graph = make(map[string]models2.GraphArray)
	c := e.newEntityCollection()
	findOptions := options.Find().SetProjection(bson.M{"title": 1, "links": 1, "categories": 1})
	cursor, err := c.Find(mongodb_official.Context, bson.D{}, findOptions)
	if err != nil {
		return
	}
	// iterate through all documents and map to graph array
	for cursor.Next(mongodb_official.Context) {
		var entity models.Entity

		// Decode the document
		if err = cursor.Decode(&entity); err != nil {
			log.Println("cursor.Decode ERROR:", err)
			return
		}
		var links []string
		for _, link := range entity.Links {
			links = append(links, link.Title)
		}
		graph[entity.GetTitle()] = models2.GraphArray{Title: entity.GetTitle(), Categories: entity.Categories, Links: links}
	}
	return
}
