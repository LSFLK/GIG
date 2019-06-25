package api

import (
	"GIG/app/models"
	"GIG/app/utility/request_handlers"
)

func (t *TestAPI) TestThatCreateEntitiesAPIWorks() {
	entity := models.Entity{}
	entity.Title = "Sri Lanka"

	var entities []models.Entity
	entities = append(entities, entity)

	//save to db
	result, _ := request_handlers.PostRequest(apiUrl+"add-batch", entities)
	t.AssertNotEqual(result, "")
}