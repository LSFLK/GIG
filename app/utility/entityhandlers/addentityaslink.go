package entityhandlers

import "GIG/app/models"

func AddEntityAsLink(entity models.Entity, linkEntity models.Entity) (models.Entity, string, error) {
	createdLinkEntity, linkEntityCreateError := CreateEntity(linkEntity)
	if linkEntityCreateError != nil {
		return entity, "", linkEntityCreateError
	}
	refVal := createdLinkEntity.ID.Hex()
	entity = entity.AddLink(refVal)
	return entity, refVal, nil
}

