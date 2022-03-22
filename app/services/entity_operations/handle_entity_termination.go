package entity_operations

import (
	"GIG/app/repositories"
	"github.com/lsflk/gig-sdk/models"
	"log"
)

func HandleEntityTermination(entity models.Entity) {
	if entity.GetTitle() != "" {
		existingEntity, err := repositories.EntityRepository{}.GetEntityBy("title", entity.GetTitle())
		if err != nil {
			log.Println(err)
		}
		err = repositories.EntityRepository{}.TerminateEntity(existingEntity, entity.GetSource(), entity.GetSourceDate())
		if err != nil {
			log.Println(err)
		}
	}

	entities, err := repositories.EntityRepository{}.GetEntities(entity.GetTitle(), entity.GetCategories(), 0, 0)
	if err != nil {
		log.Println(err)
	}

	for _, result := range entities {
		err = repositories.EntityRepository{}.TerminateEntity(result, entity.GetSource(), entity.GetSourceDate())
		if err != nil {
			log.Println(err)
		}
	}
}
