package repositories

import (
	"GIG/app/constants/mongo_drivers"
	mongodb "GIG/app/repositories/drivers/mongodb"
	mongodb_official "GIG/app/repositories/drivers/mongodb_official"
	"GIG/app/repositories/interfaces"
	"log"

	"github.com/revel/revel"
)

var repositoryHandler struct {
	entityRepository         interfaces.EntityRepositoryInterface
	userRepository           interfaces.UserRepositoryInterface
	statRepository           interfaces.StatRepositoryInterface
	normalizedNameRepository interfaces.NormalizedNameRepositoryInterface
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
