package entity_handlers

import (
	"GIG/app/models"
	"GIG/scripts/entity_handlers"
)

func (t *TestEntityHandlers) TestThatCreateEntityWorks() {
	initialEntity := models.Entity{Title: "Sri Lanka"}
	entity, _ := entity_handlers.CreateEntity(initialEntity)
	t.AssertEqual(entity.Title, "Sri Lanka")
}