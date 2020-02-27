package repositories

import (
	"GIG/app/models"
	"gopkg.in/mgo.v2/bson"
)

type iNormalizedNameRepository interface {
	AddNormalizedName(m models.NormalizedName) (normalizedName models.NormalizedName, err error)
	GetNormalizedNames() ([]models.NormalizedName, error)
	GetNormalizedName(id bson.ObjectId) (models.NormalizedName, error)
	GetNormalizedNameBy(attribute string, value string) (models.NormalizedName, error)
}

type NormalizedNameRepository struct {
}

// AddNormalizedName insert a new NormalizedName into database and returns
// last inserted normalized_name on success.
func (n NormalizedNameRepository) AddNormalizedName(m models.NormalizedName) (normalizedName models.NormalizedName, err error) {
	return repositoryHandler.normalizedNameRepository.AddNormalizedName(m)
}

// GetNormalizedNames Get all NormalizedNames from database and returns
// list of NormalizedName on success
func (n NormalizedNameRepository) GetNormalizedNames() ([]models.NormalizedName, error) {
	return repositoryHandler.normalizedNameRepository.GetNormalizedNames()
}

// GetNormalizedName Get a NormalizedName from database and returns
// a NormalizedName on success
func (n NormalizedNameRepository) GetNormalizedName(id bson.ObjectId) (models.NormalizedName, error) {
	return repositoryHandler.normalizedNameRepository.GetNormalizedName(id)
}

/**
GetEntity Get a Entity from database and returns
a models.Entity on success
 */
func (n NormalizedNameRepository) GetNormalizedNameBy(attribute string, value string) (models.NormalizedName, error) {
	return repositoryHandler.normalizedNameRepository.GetNormalizedNameBy(attribute, value)
}
