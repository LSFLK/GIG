package entity_handlers

import (
	"GIG/app/models"
	"GIG/app/utility/entity_handlers"
)

func (t *TestEntityHandlers) TestThatCreateEntitiesWorks() {
	initialEntity := models.Entity{Title: "Sri Lanka"}
	entities, _ := entity_handlers.CreateEntities(append([]models.Entity{}, initialEntity))
	t.AssertEqual(entities[0].Title, "Sri Lanka")
}
