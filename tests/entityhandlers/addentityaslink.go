package entityhandlers

import (
	"GIG/app/models"
	"GIG/app/repository"
	"GIG/app/utility/entityhandlers"
)

func (t *TestEntityHandlers) TestThatAddEntityAsLinkWorks() {
	linkEntity := models.Entity{Title: "Sri Lanka"}
	entity := models.Entity{Title: "test entity"}
	entity,_, _ = entityhandlers.AddEntityAsLink(entity, linkEntity)
	entity = repository.EagerLoad(entity)
	t.AssertEqual(entity.Links[0], "Sri Lanka")

}