package repositories

import (
	"GIG/app/constants/mongo_drivers"
	"GIG/app/repositories/mongodb"
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
		repositoryHandler.entityRepository = mongodb.EntityRepository{}
		repositoryHandler.userRepository = mongodb.UserRepository{}
		repositoryHandler.statRepository = mongodb.StatRepository{}
		repositoryHandler.normalizedNameRepository = mongodb.NormalizedNameRepository{}
	}

}
