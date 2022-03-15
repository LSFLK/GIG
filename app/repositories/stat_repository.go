package repositories

import (
	"GIG-SDK/models"
	"log"
)

type iStatRepository interface {
	AddStat(stat models.EntityStats) (models.EntityStats, error)
	GetLastStat() (models.EntityStats, error)
}

type StatRepository struct {
	iStatRepository
}

/*
AddStat insert a new Stat into database and returns
last inserted stat on success.
 */
func (e StatRepository) AddStat(stat models.EntityStats) (models.EntityStats, error) {
	log.Println("creating new stat", stat)
	stat, err := repositoryHandler.statRepository.AddStat(stat)
	return stat, err
}

/**
GetLastStat Get a Last Stat from database and returns
a models. Stat on success
 */
func (e StatRepository) GetLastStat() (models.EntityStats, error) {
	log.Println("request stat")
	stat, err := repositoryHandler.statRepository.GetLastStat()
	return stat, err
}