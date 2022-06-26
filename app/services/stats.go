package services

import (
	"GIG/app/repositories"
	"github.com/lsflk/gig-sdk/models"
	"log"
	"time"
)

func GetGraphStats(force bool) (models.EntityStats, error) {

	if force {
		// if force is true, query new and return
		newStat, err := repositories.EntityRepository{}.GetStats()
		if err != nil {
			return newStat, err
		}

		// asynchronously save new stat to stat collection
		go func(statDoc models.EntityStats) {
			_, statSaveErr := repositories.StatRepository{}.AddStat(statDoc)
			if statSaveErr != nil {
				log.Println("error saving new stats to db")
			}
		}(newStat)

		return newStat, nil
	}

	/*
		get the latest stats from db
		if stats in db are expired get new stats and save to db
		return last stats
	*/
	lastStat, err := repositories.StatRepository{}.GetLastStat()
	today := time.Now()
	expirationTime := today.Add(-1 * time.Hour)

	// entity stats are notfound in db in a recent time then generate new stats
	if err == nil && lastStat.CreatedAt.Before(expirationTime) {
		// asynchronously save new stat to stat collection
		go func() {
			// if recent stats were not found in the db or force is true, query new and return
			newStat, err := repositories.EntityRepository{}.GetStats()
			if err != nil {
				log.Println("error saving new stats to db")
			}

			_, statSaveErr := repositories.StatRepository{}.AddStat(newStat)
			if statSaveErr != nil {
				log.Println("error saving new stats to db")
			}
		}()
	}
	return lastStat, nil

}
