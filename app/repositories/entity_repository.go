package repositories

import (
	"GIG/app/models"
	"GIG/app/models/ValueType"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
)

type iEntityRepository interface {
	AddEntity(e models.Entity) (models.Entity, error)
	UpdateEntity(e models.Entity) error
	GetRelatedEntities(entity models.Entity, limit int) ([]models.Entity, error)
	GetEntities(search string, categories []string, limit int) ([]models.Entity, error)
	GetEntity(id bson.ObjectId) (models.Entity, error)
	GetEntityBy(attribute string, value string) (models.Entity, error)
	GetEntityByPreviousState(title string, date time.Time) ([]models.Entity, error)
}

type EntityRepository struct {
	iEntityRepository
}

/*
AddEntity insert a new Entity into database and returns
the entity
 */
func (e EntityRepository) AddEntity(entity models.Entity) (models.Entity, error) {
	entity = entity.SetSnippet()
	existingEntity := models.Entity{}

	/**
	get entities containing title, select the entity matching the source date
		for each value matching the title. get the most recent date that is older than source date
			iterate each entity
				iterate each titles value
					if the value is the most recent then set the corresponding entity
	 */
	var mostRecentDate time.Time
	entitiesWithMatchingTitleAndDate, _ := e.GetEntityByPreviousState(entity.GetTitle(), entity.GetSourceDate())

	// select the matching entity the given source date
	for _, resultEntity := range entitiesWithMatchingTitleAndDate {
		if resultAttribute, err := resultEntity.GetAttribute("titles"); err == nil {
			for _, resultValue := range resultAttribute.GetValues() {
				/**
				if titles match, if the source date is newer than title set date, source date is newer than mostrecentdate
				 */
				if resultValue.GetValueString() == entity.GetTitle() &&
					(resultValue.GetDate().Equal(entity.GetSourceDate()) || resultValue.GetDate().Before(entity.GetSourceDate())) &&
					mostRecentDate.Before(resultValue.GetDate()) {
					mostRecentDate = resultValue.GetDate()
					existingEntity = resultEntity
				}
			}
		}
	}

	//if an entity exists
	if existingEntity.GetTitle() != "" {
		//if the entity has a "new_title" attribute use it to change the entity title
		newTitleAttribute, err := entity.GetAttribute("new_title")
		entityIsTerminated := strings.Contains(existingEntity.GetTitle(), " - Terminated on ")
		lastTitleValue, _ := existingEntity.GetAttribute("titles")

		newTitleIsWithinLifeTime := (!entityIsTerminated) || (entityIsTerminated &&
			!newTitleAttribute.GetValue().Date.IsZero() &&
			newTitleAttribute.GetValue().Date.Before(lastTitleValue.GetValue().Date))

		newEntityIsWithinLifeTimeOfExistingEntity := (!entityIsTerminated) || (entityIsTerminated &&
			!entity.SourceDate.IsZero() &&
			entity.SourceDate.Before(lastTitleValue.GetValue().Date))

		if err == nil && newTitleIsWithinLifeTime {
			//add new title only if the new title date is before the date entity is terminated, else give an error
			fmt.Println("entity title modification found.", existingEntity.GetTitle(), "->", newTitleAttribute.GetValue().GetValueString())
			existingEntity = existingEntity.SetTitle(newTitleAttribute.GetValue())
		} else if err == nil && !newTitleIsWithinLifeTime {
			fmt.Println("new title cannot be assigned to a date after termination of the entity.")
		}

		if newEntityIsWithinLifeTimeOfExistingEntity {

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
			fmt.Println("entity exists. updated", existingEntity.GetTitle())
			return existingEntity, repositoryHandler.entityRepository.UpdateEntity(existingEntity)
		}
	}

	// if no entity exist
	entity = entity.NewEntity().SetTitle(models.Value{}.
		SetType(ValueType.String).
		SetValueString(entity.GetTitle()).
		SetDate(entity.GetSourceDate()).
		SetSource(entity.GetSource()))

	fmt.Println("creating new entity", entity.GetTitle())
	return repositoryHandler.entityRepository.AddEntity(entity)

}

/**
GetEntities Get all Entities where a given title is linked from
list of models.Entity on success
 */
func (e EntityRepository) GetRelatedEntities(entity models.Entity, limit int) ([]models.Entity, error) {
	return repositoryHandler.entityRepository.GetRelatedEntities(entity, limit)
}

/**
GetEntities Get all Entities from database and returns
list of models.Entity on success
 */
func (e EntityRepository) GetEntities(search string, categories []string, limit int) ([]models.Entity, error) {
	return repositoryHandler.entityRepository.GetEntities(search, categories, limit)
}

/**
GetEntity Get a Entity from database and returns
a models. Entity on success
 */
func (e EntityRepository) GetEntity(id bson.ObjectId) (models.Entity, error) {
	return repositoryHandler.entityRepository.GetEntity(id)
}

/**
GetEntity Get a Entity from database and returns
a models.Entity on success
 */
func (e EntityRepository) GetEntityBy(attribute string, value string) (models.Entity, error) {
	return repositoryHandler.entityRepository.GetEntityBy(attribute, value)
}

func (e EntityRepository) GetEntityByPreviousState(title string, date time.Time) ([]models.Entity, error) {
	return repositoryHandler.entityRepository.GetEntityByPreviousState(title, date)
}
