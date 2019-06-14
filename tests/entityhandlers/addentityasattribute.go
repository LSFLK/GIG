package entityhandlers

import (
	"GIG/app/models"
	"GIG/app/repository"
	"GIG/app/utility/entityhandlers"
	"github.com/revel/revel/testing"
)

type AddEntityAsAttributeTest struct {
	testing.TestSuite
}

func (t *AddEntityAsAttributeTest) Before() {
	println("Set up")
}

func (t *AddEntityAsAttributeTest) TestThatAddEntityAsAttributeWorks() {
	attributeEntity := models.Entity{Title: "Sri Lanka"}
	entity := models.Entity{Title: "test entity"}
	entity, _ = entityhandlers.AddEntityAsAttribute(entity, "testAttribute", attributeEntity)
	entity = repository.EagerLoad(entity)
	t.AssertEqual(entity.Attributes[0].Values[0].RawValue, "Sri Lanka")

}

func (t *AddEntityAsAttributeTest) After() {
	println("Tear down")
}
