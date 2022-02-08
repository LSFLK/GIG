package repositories

import (
	"GIG/app/repositories/mongodb"
)

var repositoryHandler struct{
	entityRepository iEntityRepository
	userRepository iUserRepository
	normalizedNameRepository iNormalizedNameRepository
}


func LoadRepositoryHandler() {
	repositoryHandler.entityRepository = mongodb.EntityRepository{} //change storage handler
	repositoryHandler.userRepository = mongodb.UserRepository{} //change storage handler
	repositoryHandler.normalizedNameRepository = mongodb.NormalizedNameRepository{} //change storage handler
}