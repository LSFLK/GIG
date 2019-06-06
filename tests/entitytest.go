package tests

import (
	"GIG/app/models"
	"GIG/app/utility/requesthandlers"
	"github.com/revel/revel/testing"
)

var apiUrl = "http://localhost:9000/api/add"

type EntityTest struct {
	testing.TestSuite
}

func (t *EntityTest) Before() {
	println("Set up")
}

func (t *EntityTest) TestThatSearchApiWorks() {
	t.Get("/api/search?query=Sri%20Lanka")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *EntityTest) TestThatGetEntityApiWorks() {
	t.Get("/api/get/Sri%20Lanka")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *EntityTest) TestThatCreateEntityApiWorks() {
	entity := models.Entity{}
	entity.Title = "Sri Lanka"
	entity.SourceID = "wikiAPI26750.000000"

	//save to db
	result, _ := requesthandlers.PostRequest(apiUrl, entity)
	t.AssertEqual(result.Status,"400 Bad Request")
}

func (t *EntityTest) After() {
	println("Tear down")
}
