package mongodb_official

import (
	"GIG/app/constants/database"
	"GIG/app/databases/mongodb_official"
	"github.com/lsflk/gig-sdk/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type StatRepository struct {
}

func (e StatRepository) newStatCollection() *mongodb_official.Collection {
	return mongodb_official.NewCollectionSession(database.StatCollection)
}

/*
AddStat - insert a new Stat into database and returns
last inserted stat on success.
*/
func (e StatRepository) AddStat(stat models.EntityStats) (models.EntityStats, error) {
	c := e.newStatCollection()
	defer c.Close()
	stat.CreatedAt = time.Now()
	_, err := c.Collection.InsertOne(mongodb_official.Context, stat)
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
	defer c.Close()
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"created_at", -1}}).SetLimit(1)
	err = c.Collection.FindOne(mongodb_official.Context, bson.M{}).Decode(&stat)
	return stat, err
}
