package repositories

import (
	"GIG/app/models"
	"GIG/app/models/ValueType"
	"GIG/app/repositories/mongodb"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var RepositoryHandler IHandler

type IHandler interface {
	AddEntity(e models.Entity) (models.Entity, error)
	UpdateEntity(e models.Entity) error
	GetRelatedEntities(entity models.Entity, limit int) ([]models.Entity, error)
	GetEntities(search string, categories []string, limit int) ([]models.Entity, error)
	GetEntity(id bson.ObjectId) (models.Entity, error)
	GetEntityBy(attribute string, value string) (models.Entity, error)
}

func LoadRepositoryHandler() {
	RepositoryHandler = mongodb.Repository{} //change storage handler
}

/*
AddEntity insert a new Entity into database and returns
the entity
 */
func AddEntity(entity models.Entity) (models.Entity, error) {
	existingEntity, _ := GetEntityBy("title", entity.GetTitle())

	entity = entity.SetSnippet()

	//if an entity exists
	if entity.IsEqualTo(existingEntity) {
		//if the entity has a "new_title" attribute use it to change the entity title
		newTitleAttribute, err := entity.GetAttribute("new_title")

		if err == nil { // has new_title attribute
			fmt.Println("entity title modification found.", existingEntity.GetTitle(), "->", newTitleAttribute.GetValue().GetValueString())
			existingEntity = existingEntity.SetTitle(newTitleAttribute.GetValue())
		}

		// merge links
		existingEntity = existingEntity.AddLinks(entity.GetLinks())
		// merge categories
		existingEntity = existingEntity.AddCategories(entity.GetCategories())
		// merge attributes

		for name := range entity.GetAttributes() {
			if name != "new_title" && name != "title" {
				entityAttribute, _ := entity.GetAttribute(name)
				existingEntity = existingEntity.SetAttribute(name, entityAttribute.GetValue())
			}
		}
		fmt.Println("entity exists. updated", entity.GetTitle())

		return existingEntity, RepositoryHandler.UpdateEntity(existingEntity)
	} else {
		// if no entity exist
		entity := entity.NewEntity().SetTitle(models.Value{}.
			SetType(ValueType.String).
			SetValueString(entity.GetTitle()).
			SetDate(time.Now()).
			SetSource(entity.GetSource()))

		fmt.Println("creating new entity", entity.GetTitle())
		return RepositoryHandler.AddEntity(entity)
	}

}

/**
GetEntities Get all Entities where a given title is linked from
list of models.Entity on success
 */
func GetRelatedEntities(entity models.Entity, limit int) ([]models.Entity, error) {
	return RepositoryHandler.GetRelatedEntities(entity, limit)
}

/**
GetEntities Get all Entities from database and returns
list of models.Entity on success
 */
func GetEntities(search string, categories []string, limit int) ([]models.Entity, error) {
	return RepositoryHandler.GetEntities(search, categories, limit)
}

/**
GetEntity Get a Entity from database and returns
a models. Entity on success
 */
func GetEntity(id bson.ObjectId) (models.Entity, error) {
	return RepositoryHandler.GetEntity(id)
}

/**
GetEntity Get a Entity from database and returns
a models.Entity on success
 */
func GetEntityBy(attribute string, value string) (models.Entity, error) {
	return RepositoryHandler.GetEntityBy(attribute, value)
}
