package entityhandlers

import (
	"GIG/app/models"
	"GIG/app/utility/entityhandlers"
)

func (t *TestEntityHandlers) TestThatCreateEntityWorks() {
	initialEntity := models.Entity{Title: "Sri Lanka"}
	entity, _ := entityhandlers.CreateEntity(initialEntity)
	t.AssertEqual(entity.Title, "Sri Lanka")

}