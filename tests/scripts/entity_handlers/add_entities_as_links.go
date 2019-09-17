package entity_handlers

import (
	"GIG/app/models"
	"GIG/scripts/entity_handlers"
)

func (t *TestEntityHandlers) TestThatAddEntitiesAsLinksWorks() {
	linkEntity := models.Entity{Title: "Sri Lanka"}

	entity := models.Entity{Title: "test entity"}
	entity, _ = entity_handlers.AddEntitiesAsLinks(entity, append([]models.Entity{}, linkEntity))
	t.AssertEqual(entity.Links[0], "Sri Lanka")

}
