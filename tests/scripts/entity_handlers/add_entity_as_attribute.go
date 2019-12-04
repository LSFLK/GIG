package entity_handlers

import (
	"GIG/app/models"
	"GIG/app/repositories/mongodb"
	"GIG/scripts/entity_handlers"
	"fmt"
)

func (t *TestEntityHandlers) TestThatAddEntityAsAttributeWorks() {
	attributeEntity, err := mongodb.GetEntityBy("title", "Sri Lanka")
	if err != nil {
		t.AssertNotFound()
	}
	entity := models.Entity{Title: "test entity"}
	entity, attributeEntity, _ = entity_handlers.AddEntityAsAttribute(entity, "testAttribute", attributeEntity)
	fmt.Println(attributeEntity.ID)
	fmt.Println(entity.Attributes[0].Values)
	t.AssertEqual(entity.Attributes[0].Values[0].RawValue, attributeEntity.ID.Hex())

}