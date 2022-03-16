package managers

import (
	"GIG-SDK/models"
	"log"
)

func (e EntityManager) MergeEntityTitle(existingEntity models.Entity, newEntity models.Entity) (models.Entity, error) {
	//if the entity has a "new_title" attribute use it to change the entity title
	newTitleAttribute, err := newEntity.GetAttribute("new_title")
	lastTitleAttribute, _ := existingEntity.GetAttribute("titles")

	isValidTitle := e.NewEntityTitleIsWithinLifetimeOfExistingEntity(newTitleAttribute, lastTitleAttribute, existingEntity.IsTerminated())

	if err == nil && isValidTitle {
		//add new title only if the new title date is before the date entity is terminated, else give an error
		log.Println("entity title modification found.", existingEntity.GetTitle(), "->", newTitleAttribute.GetValue().GetValueString())
		existingEntity = existingEntity.SetTitle(newTitleAttribute.GetValue())
	} else if err == nil && !isValidTitle {
		log.Println("new title cannot be assigned to a date after termination of the entity.")
	}
	return existingEntity, err
}
