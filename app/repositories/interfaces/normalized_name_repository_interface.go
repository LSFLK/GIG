package interfaces

import "github.com/lsflk/gig-sdk/models"

type NormalizedNameRepositoryInterface interface {
	AddNormalizedName(m models.NormalizedName) (normalizedName models.NormalizedName, err error)
	GetNormalizedNames(searchString string, limit int) ([]models.NormalizedName, error)
	GetNormalizedName(id string) (models.NormalizedName, error)
	GetNormalizedNameBy(attribute string, value string) (models.NormalizedName, error)
}
