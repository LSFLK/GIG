package repositories

import (
	"GIG/app/models"
	"GIG/app/models/ValueType"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
)

type iEntityRepository interface {
	AddEntity(e models.Entity) (models.Entity, error)
	UpdateEntity(e models.Entity) error
	GetRelatedEntities(entity models.Entity, limit int, offset int) ([]models.Entity, error)
	GetEntities(search string, categories []string, limit int, offset int) ([]models.Entity, error)
	GetEntity(id bson.ObjectId) (models.Entity, error)
	GetEntityBy(attribute string, value string) (models.Entity, error)
	GetEntityByPreviousState(title string, date time.Time) ([]models.Entity, error)
	DeleteEntity(entity models.Entity) error
}

type EntityRepository struct {
	iEntityRepository
}

/*
AddEntity insert a new Entity into database and returns
the entity
 */
func (e EntityRepository) AddEntity(entity models.Entity) (models.Entity, int, error) {
	if strings.TrimSpace(entity.GetTitle()) == "" {
		return entity, 406, errors.New("title cannot be empty")
	}
	normalizedTitle := ""
	entity = entity.SetSnippet()
	if isFromVerifiedSource(entity) {
		AddTitleToNormalizationDatabase(entity.GetTitle(), entity.GetTitle())
	} else {
		entityTitle, normalizationErr := NormalizeEntityTitle(entity.GetTitle())
		if normalizationErr == nil {
			normalizedTitle = entityTitle
		} else {
			entity = entity.AddCategory("arbitrary-entities")
		}
	}
	var (
		existingEntity models.Entity
		err            error
	)

	if entity.GetSourceDate().IsZero() {
		existingEntity, err = e.GetEntityBy("title", entity.GetTitle())
	} else {
		existingEntity, err = e.GetEntityByPreviousTitle(entity.GetTitle(), entity.GetSourceDate())
	}

	if entityIsCompatible, existingEntity := CheckEntityCompatibility(existingEntity, entity); entityIsCompatible && err == nil {

		if existingEntity.GetImageURL() == "" {
			existingEntity = existingEntity.SetImageURL(entity.GetImageURL())
		}
		if existingEntity.GetSource() == "" {
			existingEntity = existingEntity.SetSource(entity.GetSource())
		}
		if existingEntity.GetSourceSignature() == "" {
			existingEntity = existingEntity.SetSourceSignature(entity.GetSourceSignature())
		}

		fmt.Println("entity exists. updating", existingEntity.GetTitle())
		return existingEntity, 202, repositoryHandler.entityRepository.UpdateEntity(existingEntity)
	}

	titleValue := models.Value{}.
		SetType(ValueType.String).
		SetValueString(entity.GetTitle()).
		SetDate(entity.GetSourceDate()).
		SetSource(entity.GetSource())

	// if no entity exist
	entity = entity.NewEntity().SetTitle(titleValue)
	if normalizedTitle != "" {
		entity = entity.NewEntity().
			SetTitle(titleValue.
				SetValueString(normalizedTitle).
				SetSource("normalizer"))
	}

	fmt.Println("creating new entity", entity.GetTitle())
	existingEntity, err = repositoryHandler.entityRepository.AddEntity(entity)
	return existingEntity, 201, err

}

/**
GetEntities Get all Entities where a given title is linked from
list of models.Entity on success
 */
func (e EntityRepository) GetRelatedEntities(entity models.Entity, limit int, offset int) ([]models.Entity, error) {
	return repositoryHandler.entityRepository.GetRelatedEntities(entity, limit, offset)
}

/**
GetEntities Get all Entities from database and returns
list of models.Entity on success
 */
func (e EntityRepository) GetEntities(search string, categories []string, limit int, offset int) ([]models.Entity, error) {
	return repositoryHandler.entityRepository.GetEntities(search, categories, limit, offset)
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

func (e EntityRepository) GetEntityByPreviousTitle(title string, searchDate time.Time) (models.Entity, error) {
	/**
	get entities containing title, select the entity matching the source date
		for each value matching the title. get the most recent date that is older than source date
			iterate each entity
				iterate each titles value
					if the value is the most recent then set the corresponding entity
	 */
	var mostRecentDate time.Time
	entitiesWithMatchingTitleAndDate, err := repositoryHandler.entityRepository.GetEntityByPreviousState(title, searchDate)
	existingEntity := models.Entity{}
	if err != nil {
		return existingEntity, err
	}

	// select the matching entity the given source date
	for _, resultEntity := range entitiesWithMatchingTitleAndDate {
		if resultAttribute, err := resultEntity.GetAttribute("titles"); err == nil {
			resultValue := resultAttribute.GetValueByDate(searchDate)
			resultValue2 := resultAttribute.GetValueByDate(searchDate.Add(time.Duration(-1) * time.Second))
			/**
				if titles match, if the source date is newer than title set date, source date is newer than most recent date
				 */
			if resultValue.GetValueString() == title && mostRecentDate.Before(resultValue.GetDate()) {
				mostRecentDate = resultValue.GetDate()
				existingEntity = resultEntity
			}
			if resultValue2.GetValueString() == title && mostRecentDate.Before(resultValue2.GetDate()) {
				mostRecentDate = resultValue2.GetDate()
				existingEntity = resultEntity
			}
			if (resultValue.GetValueString() == title && mostRecentDate.IsZero() && resultValue.GetDate().IsZero()) ||
				(resultValue2.GetValueString() == title && mostRecentDate.IsZero() && resultValue2.GetDate().IsZero()) {
				existingEntity = resultEntity
			}
		}
	}
	if existingEntity.GetTitle() == "" {
		return existingEntity, errors.New("no matching entity found")
	}
	return existingEntity, nil
}

func (e EntityRepository) TerminateEntity(existingEntity models.Entity, sourceString string, terminationDate time.Time) error {
	if !existingEntity.IsTerminated() && existingEntity.GetSourceDate().Before(terminationDate) {
		entity := existingEntity.
			SetAttribute("lifeStatus",
				models.Value{
					ValueType:   "string",
					ValueString: "Terminated",
					Source:      sourceString,
					Date:        terminationDate,
					UpdatedAt:   time.Now(),
				}).SetAttribute("new_title",
			models.Value{
				ValueType:   "string",
				ValueString: existingEntity.GetTitle() + " - Terminated on " + terminationDate.Format("2006-01-02"),
				Source:      sourceString,
				Date:        terminationDate,
				UpdatedAt:   time.Now(),
			})
		//save to db
		if entityIsCompatible, existingEntity := CheckEntityCompatibility(existingEntity, entity); entityIsCompatible {
			existingEntity = existingEntity.RemoveAttribute("new_title")
			fmt.Println("entity exists. terminating", existingEntity.GetTitle())
			return repositoryHandler.entityRepository.UpdateEntity(existingEntity)
		}
	}
	return nil
}

func (e EntityRepository) DeleteEntity(entity models.Entity) error {
	return repositoryHandler.entityRepository.DeleteEntity(entity)
}
