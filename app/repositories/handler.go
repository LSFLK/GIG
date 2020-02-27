package repositories

import (
	"GIG/app/repositories/mongodb"
)

var repositoryHandler struct{
	entityRepository iEntityRepository
	normalizedNameRepository iNormalizedNameRepository
}


func LoadRepositoryHandler() {
	repositoryHandler.entityRepository = mongodb.EntityRepository{} //change storage handler
	repositoryHandler.normalizedNameRepository = mongodb.NormalizedNameRepository{} //change storage handler
}