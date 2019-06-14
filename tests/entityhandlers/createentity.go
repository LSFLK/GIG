package entityhandlers

import (
	"GIG/app/models"
	"GIG/app/utility/entityhandlers"
	"github.com/revel/revel/testing"
)

type CreateEntityTest struct {
	testing.TestSuite
}

func (t *CreateEntityTest) Before() {
	println("Set up")
}

func (t *CreateEntityTest) TestThatCreateEntityWorks() {
	initialEntity := models.Entity{Title: "Sri Lanka"}
	entity, _ := entityhandlers.CreateEntity(initialEntity)
	t.AssertEqual(entity.Title, "Sri Lanka")

}

func (t *CreateEntityTest) After() {
	println("Tear down")
}
