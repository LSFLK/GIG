package api

import (
	"GIG/app/models"
	"GIG/commons/request_handlers"
)

func (t *TestAPI) TestThatCreateEntityAPIWorks() {
	entity := models.Entity{}
	entity.Title = "Sri Lanka"

	//save to db
	result, _ := request_handlers.PostRequest(apiUrl+"add", entity)
	t.AssertNotEqual(result, "")
}