package databases

import (
	"GIG/app/databases/mongodb"
)

func LoadDatabaseHandler() {
	mongodb.LoadMongo()		// change database config loader
}
