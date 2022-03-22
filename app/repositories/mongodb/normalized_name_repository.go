package mongodb

import (
	"GIG/app/databases/mongodb"
	"github.com/lsflk/gig-sdk/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type NormalizedNameRepository struct {
}

func (n NormalizedNameRepository) newNormalizedNameCollection() *mongodb.Collection {
	c := mongodb.NewCollectionSession("normalized_names")
	textIndex := mgo.Index{
		Key: []string{"$text:search_text"},
		Weights: map[string]int{
			"search_text": 1,
		},
		Name:   "textIndex",
		Unique: true,
	}
	searchTextIndex := mgo.Index{
		Key:    []string{"search_text"},
		Name:   "searchTextIndex",
		Unique: true,
	}
	c.Session.EnsureIndex(textIndex)
	c.Session.EnsureIndex(searchTextIndex)
	return c
}

// AddNormalizedName insert a new NormalizedName into database and returns
// last inserted normalized_name on success.
func (n NormalizedNameRepository) AddNormalizedName(m models.NormalizedName) (normalizedName models.NormalizedName, err error) {
	c := n.newNormalizedNameCollection()
	defer c.Close()
	m = m.NewNormalizedName()
	return m, c.Session.Insert(m)
}

// GetNormalizedNames Get all NormalizedNames from database and returns
// list of NormalizedName on success
func (n NormalizedNameRepository) GetNormalizedNames(searchString string, limit int) ([]models.NormalizedName, error) {
	var (
		normalizedNames []models.NormalizedName
		err             error
		resultQuery     *mgo.Query
	)

	query := bson.M{}
	c := n.newNormalizedNameCollection()
	defer c.Close()

	if searchString != "" {
		query = bson.M{
			"$text": bson.M{"$search": searchString},
		}
	}

	resultQuery = c.Session.Find(query).Select(bson.M{
		"score": bson.M{"$meta": "textScore"}}).Sort("$textScore:score")

	err = resultQuery.Limit(limit).All(&normalizedNames)

	return normalizedNames, err
}

// GetNormalizedName Get a NormalizedName from database and returns
// a NormalizedName on success
func (n NormalizedNameRepository) GetNormalizedName(id bson.ObjectId) (models.NormalizedName, error) {
	var (
		normalizedName models.NormalizedName
		err            error
	)

	c := n.newNormalizedNameCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"_id": id}).One(&normalizedName)
	return normalizedName, err
}

/**
GetEntity Get a Entity from database and returns
a models.Entity on success
 */
func (n NormalizedNameRepository) GetNormalizedNameBy(attribute string, value string) (models.NormalizedName, error) {
	var (
		normalizedName models.NormalizedName
		err            error
	)

	c := n.newNormalizedNameCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{attribute: value}).One(&normalizedName)
	return normalizedName, err
}
