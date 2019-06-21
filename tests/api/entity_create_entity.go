package api

import (
	"GIG/app/models"
	"GIG/app/utility/request_handlers"
)

func (t *TestAPI) TestThatCreateEntityAPIWorks() {
	entity := models.Entity{}
	entity.Title = "Sri Lanka"

	//save to db
	result, _ := request_handlers.PostRequest(apiUrl+"add", entity)
	defer result.Body.Close()
	t.AssertEqual(result.Status, "201 Created")
}