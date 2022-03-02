package services

import (
	"GIG/app/repositories"
	"log"
)

func GetGraphStats(){
	/*
	get number of entities
	get list of categories
	get number of relations formed
	 */
	 entityStats, err:= repositories.EntityRepository{}.GetStats()

	 if err!=nil{
	 	log.Println("error reading entity stats:", err)
	 }
	 log.Println(entityStats)
}
