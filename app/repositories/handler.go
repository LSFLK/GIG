package repositories

import (
	"GIG/app/constants/mongo_drivers"
	mongodb2 "GIG/app/repositories/drivers/mongodb"
	mongodb_official2 "GIG/app/repositories/drivers/mongodb_official"
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
		repositoryHandler.entityRepository = mongodb2.EntityRepository{}
		repositoryHandler.userRepository = mongodb2.UserRepository{}
		repositoryHandler.statRepository = mongodb2.StatRepository{}
		repositoryHandler.normalizedNameRepository = mongodb2.NormalizedNameRepository{}
	case mongo_drivers.MongodbOfficial:
		repositoryHandler.entityRepository = mongodb_official2.EntityRepository{}
		repositoryHandler.userRepository = mongodb_official2.UserRepository{}
		repositoryHandler.statRepository = mongodb_official2.StatRepository{}
		repositoryHandler.normalizedNameRepository = mongodb_official2.NormalizedNameRepository{}
	}

}
