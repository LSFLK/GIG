package api

import (
	"GIG/app/models"
	"GIG/app/utility/request_handlers"
	"github.com/revel/revel/testing"
)

var apiUrl = "http://localhost:9000/api/add"

type EntityAPITest struct {
	testing.TestSuite
}

func (t *EntityAPITest) Before() {
	println("Set up")
}

func (t *EntityAPITest) TestThatSearchApiWorks() {
	t.Get("/api/search?query=Sri%20Lanka")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *EntityAPITest) TestThatGetEntityApiWorks() {
	t.Get("/api/get/Sri%20Lanka")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *EntityAPITest) TestThatCreateEntityApiWorks() {
	entity := models.Entity{}
	entity.Title = "Sri Lanka"

	//save to db
	result, _ := request_handlers.PostRequest(apiUrl, entity)
	defer result.Body.Close()
	t.AssertEqual(result.Status,"201 Created")
}

func (t *EntityAPITest) After() {
	println("Tear down")
}
