package mongodb

import (
	"GIG/app/databases/mongodb"
	"GIG/app/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type NormalizedNameRepository struct {
}

func (n NormalizedNameRepository) newNormalizedNameCollection() *mongodb.Collection {
	c := mongodb.NewCollectionSession("normalized_names")
	searchTextIndex := mgo.Index{
		Key:    []string{"searchText"},
		Name:   "searchTextIndex",
		Unique: true,
	}
	c.Session.EnsureIndex(searchTextIndex)
	return c
}

// AddNormalizedName insert a new NormalizedName into database and returns
// last inserted normalized_name on success.
func (n NormalizedNameRepository) AddNormalizedName(m models.NormalizedName) (normalizedName models.NormalizedName, err error) {
	c := n.newNormalizedNameCollection()
	defer c.Close()
	m.ID = bson.NewObjectId()
	m.CreatedAt = time.Now()
	return m, c.Session.Insert(m)
}

// GetNormalizedNames Get all NormalizedName from database and returns
// list of NormalizedName on success
func (n NormalizedNameRepository) GetNormalizedNames() ([]models.NormalizedName, error) {
	var (
		normalizedNames []models.NormalizedName
		err             error
	)

	c := n.newNormalizedNameCollection()
	defer c.Close()

	err = c.Session.Find(nil).Sort("-createdAt").All(&normalizedNames)
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
