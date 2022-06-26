package mongodb_official

import (
	"GIG/app/constants/database"
	"GIG/app/databases/mongodb_official"
	"GIG/app/repositories/interfaces"
	"github.com/lsflk/gig-sdk/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type StatRepository struct {
	interfaces.StatRepositoryInterface
}

func (e StatRepository) newStatCollection() *mongo.Collection {
	return mongodb_official.GetCollection(database.StatCollection)
}

/*
AddStat - insert a new Stat into database and returns
last inserted stat on success.
*/
func (e StatRepository) AddStat(stat models.EntityStats) (models.EntityStats, error) {
	c := e.newStatCollection()
	stat.CreatedAt = time.Now()
	_, err := c.InsertOne(mongodb_official.Context, stat)
	return stat, err
}

/*
GetLastStat - Get a Last Stat from database and returns
a models. Stat on success
*/
func (e StatRepository) GetLastStat() (models.EntityStats, error) {
	var (
		stat models.EntityStats
		err  error
	)

	c := e.newStatCollection()
	findOptions := options.FindOne()
	findOptions.SetSort(bson.D{{"created_at", -1}})
	err = c.FindOne(mongodb_official.Context, bson.M{}, findOptions).Decode(&stat)
	return stat, err
}
