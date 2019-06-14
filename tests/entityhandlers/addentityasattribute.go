package entityhandlers

import "GIG/app/models"

func AddEntityAsAttribute(entity models.Entity, attributeName string, attributeEntity models.Entity) (models.Entity, error) {
	createdAttributeEntity, attributeEntityCreateError := CreateEntity(attributeEntity)
	if attributeEntityCreateError != nil {
		return entity, attributeEntityCreateError
	}
	refVal := createdAttributeEntity.ID.Hex()
	entity = entity.SetAttribute(attributeName, models.Value{
		Type:     "objectId",
		RawValue: refVal,
	})
	entity = entity.AddLink(refVal)
	return entity, attributeEntityCreateError
}
