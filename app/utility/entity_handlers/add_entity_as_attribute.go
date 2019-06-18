package entity_handlers

import "GIG/app/models"

func AddEntityAsAttribute(entity models.Entity, attributeName string, attributeEntity models.Entity) (models.Entity, error) {
	entity, refVal, err := AddEntityAsLink(entity, attributeEntity)
	if err != nil {
		return entity, err
	}
	entity = entity.SetAttribute(attributeName, models.Value{
		Type:     "objectId",
		RawValue: refVal,
	})

	return entity, nil
}