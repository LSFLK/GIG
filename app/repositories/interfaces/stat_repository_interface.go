package interfaces

import "github.com/lsflk/gig-sdk/models"

type StatRepositoryInterface interface {
	AddStat(stat models.EntityStats) (models.EntityStats, error)
	GetLastStat() (models.EntityStats, error)
}
