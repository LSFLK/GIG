package repositories

import (
	"GIG/app/constants/mongo_drivers"
	"GIG/app/repositories/mongodb"
	"GIG/app/repositories/mongodb_official"
	"log"

	"github.com/revel/revel"
)

var repositoryHandler struct {
	entityRepository         iEntityRepository
	userRepository           iUserRepository
	statRepository           iStatRepository
	normalizedNameRepository iNormalizedNameRepository
}

func LoadRepositoryHandler() {
	driver, found := revel.Config.String("mongo.driver")
	if !found {
		log.Fatal("MongoDB driver not configured")
	}
	switch driver {
	case mongo_drivers.Mongodb:
		repositoryHandler.entityRepository = mongodb.EntityRepository{}
		repositoryHandler.userRepository = mongodb.UserRepository{}
		repositoryHandler.statRepository = mongodb.StatRepository{}
		repositoryHandler.normalizedNameRepository = mongodb.NormalizedNameRepository{}
	case mongo_drivers.MongodbOfficial:
		repositoryHandler.entityRepository = mongodb_official.EntityRepository{}
		repositoryHandler.userRepository = mongodb_official.UserRepository{}
		repositoryHandler.statRepository = mongodb_official.StatRepository{}
		repositoryHandler.normalizedNameRepository = mongodb_official.NormalizedNameRepository{}
	}

}
