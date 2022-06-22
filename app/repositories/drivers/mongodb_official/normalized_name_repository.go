package mongodb_official

import (
	"GIG/app/constants/database"
	"GIG/app/databases/mongodb_official"
	"GIG/app/repositories/interfaces"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	"github.com/lsflk/gig-sdk/models"
	"go.mongodb.org/mongo-driver/bson"
)

type NormalizedNameRepository struct {
	interfaces.NormalizedNameRepositoryInterface
}

func (n NormalizedNameRepository) newNormalizedNameCollection() *mongodb_official.Collection {
	return mongodb_official.NewCollectionSession(database.NormalizedNameCollection)
}

// AddNormalizedName insert a new NormalizedName into database and returns
// last inserted normalized_name on success.
func (n NormalizedNameRepository) AddNormalizedName(m models.NormalizedName) (normalizedName models.NormalizedName, err error) {
	c := n.newNormalizedNameCollection()
	defer c.Close()
	m = m.NewNormalizedName()
	_, err = c.Collection.InsertOne(mongodb_official.Context, m)
	return m, err
}

// GetNormalizedNames Get all NormalizedNames from database and returns
// list of NormalizedName on success
func (n NormalizedNameRepository) GetNormalizedNames(searchString string, limit int) ([]models.NormalizedName, error) {
	var (
		normalizedNames []models.NormalizedName
		err             error
	)

	query := bson.M{}
	c := n.newNormalizedNameCollection()
	defer c.Close()

	if searchString != "" {
		query = bson.M{
			"$text": bson.M{"$search": searchString},
		}
	}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"textScore:score", -1}}).
		SetLimit(int64(limit))
	cursor, err := c.Collection.Find(mongodb_official.Context, query, findOptions)
	if err != nil {
		return normalizedNames, err
	}
	err = cursor.All(mongodb_official.Context, &normalizedNames)
	log.Println(normalizedNames, err)
	return normalizedNames, err
}

// GetNormalizedName Get a NormalizedName from database and returns
// a NormalizedName on success
func (n NormalizedNameRepository) GetNormalizedName(id string) (models.NormalizedName, error) {
	var (
		normalizedName models.NormalizedName
		err            error
	)

	c := n.newNormalizedNameCollection()
	defer c.Close()

	cursor := c.Collection.FindOne(mongodb_official.Context, bson.M{"_id": id})
	err = cursor.Decode(&normalizedName)
	return normalizedName, err
}

/*
GetNormalizedNameBy - Get a Entity from database and returns
a models.Entity on success
*/
func (n NormalizedNameRepository) GetNormalizedNameBy(attribute string, value string) (models.NormalizedName, error) {
	var (
		normalizedName models.NormalizedName
		err            error
	)

	c := n.newNormalizedNameCollection()
	defer c.Close()

	cursor := c.Collection.FindOne(mongodb_official.Context, bson.M{attribute: value})
	err = cursor.Decode(&normalizedName)
	return normalizedName, err
}
