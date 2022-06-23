package interfaces

import (
	"github.com/lsflk/gig-sdk/models"
	"time"
)

type EntityRepositoryInterface interface {
	AddEntity(e models.Entity) (models.Entity, error)
	UpdateEntity(e models.Entity) error
	GetRelatedEntities(entity models.Entity, limit int, offset int) ([]models.Entity, error)
	GetEntities(search string, categories []string, limit int, offset int) ([]models.Entity, error)
	GetEntity(id string) (models.Entity, error)
	GetEntityBy(attribute string, value string) (models.Entity, error)
	GetEntityByPreviousTitle(title string, date time.Time) (models.Entity, error)
	DeleteEntity(entity models.Entity) error
	GetStats() (models.EntityStats, error)
}
