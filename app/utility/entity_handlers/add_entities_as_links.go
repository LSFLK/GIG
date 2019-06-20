package entity_handlers

import "GIG/app/models"

func AddEntitiesAsLinks(entity models.Entity, linkEntities []models.Entity) (models.Entity, error) {
	createdLinkEntities, linkEntityCreateError := CreateEntities(linkEntities)
	if linkEntityCreateError != nil {
		return entity, linkEntityCreateError
	}
	for _, linkEntity := range createdLinkEntities {
		entity = entity.AddLink(linkEntity)
	}
	return entity, nil
}
