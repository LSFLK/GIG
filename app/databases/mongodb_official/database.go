package mongodb_official

import "go.mongodb.org/mongo-driver/mongo"

type Database struct {
	s        *mongo.Session
	name     string
	database *mongo.Database
}

func (db *Database) Connect() {

	db.s = service.Session()
	database := *service.client.Database(db.name)
	db.database = &database

}

func newDBSession(name string) *mongo.Database {

	return service.client.Database(name)
}
