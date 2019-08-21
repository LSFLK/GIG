package entity_handlers

import (
	"GIG/app/models"
	"GIG/app/repository"
	"GIG/app/scripts/entity_handlers"
)

func (t *TestEntityHandlers) TestThatAddEntityAsLinkWorks() {
	linkEntity := models.Entity{Title: "Sri Lanka"}
	entity := models.Entity{Title: "test entity"}
	entity, _, _ = entity_handlers.AddEntityAsLink(entity, linkEntity)
	entity = repository.EagerLoad(entity)
	t.AssertEqual(entity.LoadedLinks[0].Title, "Sri Lanka")

}
