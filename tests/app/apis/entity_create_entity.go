package apis

import (
	"GIG/app/models"
	"GIG/app/models/ValueType"
	"GIG/commons/request_handlers"
	"time"
)

func (t *TestAPI) TestThatCreateEntityAPIWorks() {
	entity := models.Entity{}.SetTitle(models.Value{}.
		SetType(ValueType.String).
		SetValueString("Sri Lanka").
		SetDate(time.Now()).
		SetSource("unit test"))

	//save to db
	result, _ := request_handlers.PostRequest(apiUrl+"add", entity)
	t.AssertNotEqual(result, "")
}
