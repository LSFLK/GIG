package entity_handlers

import "GIG/app/models"

/**
Add list of related entities to a given entity
 */
func AddEntitiesAsLinks(entity models.Entity, linkEntities []models.Entity) (models.Entity, error) {
	createdLinkEntities, linkEntityCreateError := CreateEntities(linkEntities)
	if linkEntityCreateError != nil {
		return entity, linkEntityCreateError
	}
	for _, linkEntity := range createdLinkEntities {
		entity = entity.AddLink(linkEntity.Title)
	}
	return entity, nil
}
