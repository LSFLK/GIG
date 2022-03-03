package services

import (
	"GIG-SDK/models"
	"GIG/app/repositories"
	"log"
	"time"
)

func GetGraphStats(force bool) (models.EntityStats, error) {

	/*
	get latest stats from db
	if stats in db are expired get new stats and save to db
	return new stats
	 */
	 if !force { // forcefully run the query to generate fresh stats
		 lastStat, err := repositories.StatRepository{}.GetLastStat()
		 today := time.Now()
		 expirationTime := today.Add(-1 * time.Hour)

		 // entity stats found in db in a recent time then return it
		 if err == nil && lastStat.CreatedAt.After(expirationTime) {
			 log.Println("entity stat already available")
			 return lastStat, nil
		 }
	 }

	// if entity stats were not found in the db query new and return
	newStat, err := repositories.EntityRepository{}.GetStats()
	if err != nil {
		return newStat, err
	}
	return repositories.StatRepository{}.AddStat(newStat)

}
