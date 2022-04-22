package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type Database struct {
	s       *mongo.Session
	name    string
	session *mongo.Database
}

func (db *Database) Connect() {

	db.s = service.Session()
	session := *db.s.DB(db.name)
	db.session = &session

}

func newDBSession(name string) *Database {

	var db = Database{
		name: name,
	}
	db.Connect()
	return &db
}
