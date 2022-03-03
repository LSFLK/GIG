package services

import (
	"GIG-SDK/models"
	"GIG/app/repositories"
)

func GetGraphStats() (models.EntityStats, error) {

	/*
	get latest stats from db
	if stats in db are expired get new stats and save to db
	return new stats
	 */

	return repositories.EntityRepository{}.GetStats()

}
