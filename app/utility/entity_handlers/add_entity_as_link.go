package entity_handlers

import "GIG/app/models"

func AddEntityAsLink(entity models.Entity, linkEntity models.Entity) (models.Entity, models.Entity, error) {
	createdLinkEntity, linkEntityCreateError := CreateEntity(linkEntity)
	if linkEntityCreateError != nil {
		return entity, createdLinkEntity, linkEntityCreateError
	}
	entity = entity.AddLink(createdLinkEntity)
	return entity, createdLinkEntity, nil
}

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
