package services

import (
	"GIG/app/repositories"
	"github.com/lsflk/gig-sdk/models"
	"log"
	"time"
)

func GetGraphStats(force bool) (models.EntityStats, error) {

	/*
		get the latest stats from db
		if stats in db are expired get new stats and save to db
		return new stats
	*/
	if !force { // if force is false, then check if previous stored recent stat record is available
		lastStat, err := repositories.StatRepository{}.GetLastStat()
		today := time.Now()
		expirationTime := today.Add(-1 * time.Hour)

		// entity stats found in db in a recent time then return it
		if err == nil && lastStat.CreatedAt.After(expirationTime) {
			log.Println("entity stat already available")
			return lastStat, nil
		}
	}

	// if recent stats were not found in the db or force is true, query new and return
	newStat, err := repositories.EntityRepository{}.GetStats()
	if err != nil {
		return newStat, err
	}

	// asynchronously save new stat to stat collection
	go func(statDoc models.EntityStats) {
		_, statSaveErr := repositories.StatRepository{}.AddStat(statDoc)
		if statSaveErr != nil {

		}
	}(newStat)
	
	return newStat, nil

}
