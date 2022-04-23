package mongodb

import "gopkg.in/mgo.v2"

type Database struct {
	s        *mgo.Session
	name     string
	database *mgo.Database
}

func (db *Database) Connect() {

	db.s = service.Session()
	database := *db.s.DB(db.name)
	db.database = &database

}

func newDBSession(name string) *Database {

	var db = Database{
		name: name,
	}
	db.Connect()
	return &db
}
