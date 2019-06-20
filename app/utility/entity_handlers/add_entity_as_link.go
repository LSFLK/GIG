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