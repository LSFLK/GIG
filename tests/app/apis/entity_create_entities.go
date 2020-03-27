package apis

import (
	"GIG/app/models"
	"GIG/app/models/ValueType"
	"GIG/commons/request_handlers"
	"time"
)

func (t *TestAPI) TestThatCreateEntitiesAPIWorks() {
	entity := models.Entity{}.SetTitle(models.Value{}.
		SetType(ValueType.String).
		SetValueString("Sri Lanka").
		SetDate(time.Now()).
		SetSource("unit test"))

	var entities []models.Entity
	entities = append(entities, entity)

	//save to db
	result, _ := request_handlers.PostRequest(apiUrl+"add-batch", entities)
	t.AssertNotEqual(result, "")
}
