package repositories

import (
	"GIG/app/constants/error_messages"
	"GIG/app/repositories/functions"
	"GIG/app/utilities/managers"
	"errors"
	"github.com/lsflk/gig-sdk/enums/ValueType"
	"github.com/lsflk/gig-sdk/libraries"
	"github.com/lsflk/gig-sdk/models"
	"gopkg.in/mgo.v2/bson"
	"log"
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
	GetEntityByPreviousTitle(title string, date time.Time) (models.Entity, error)
	DeleteEntity(entity models.Entity) error
	GetStats() (models.EntityStats, error)
}

type EntityRepository struct {
	iEntityRepository
}

/*
AddEntity insert a new Entity into database and returns
the entity
*/
func (e EntityRepository) AddEntity(entity models.Entity) (models.Entity, error) {
	if strings.TrimSpace(entity.GetTitle()) == "" {
		return entity, errors.New("title cannot be empty")
	}
	entity = entity.SetSnippet()
	entity, normalizedTitle := e.normalizeEntity(entity)
	existingEntity, err := e.getExistingEntity(entity)
	entityIsCompatible, existingEntity := managers.EntityManager{}.CheckEntityCompatibility(existingEntity, entity)

	// if existing entity found
	if entityIsCompatible && err == nil {
		return e.updateExistingEntity(entity, existingEntity)
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

	log.Println("creating new entity", entity.GetTitle())
	existingEntity, err = repositoryHandler.entityRepository.AddEntity(entity)
	return existingEntity, err

}

/*
GetRelatedEntities - Get all Entities where a given title is linked from
list of models.Entity on success
*/
func (e EntityRepository) GetRelatedEntities(entity models.Entity, limit int, offset int) ([]models.Entity, error) {
	return repositoryHandler.entityRepository.GetRelatedEntities(entity, limit, offset)
}

/*
GetEntities - Get all Entities from database and returns
list of models.Entity on success
*/
func (e EntityRepository) GetEntities(search string, categories []string, limit int, offset int) ([]models.Entity, error) {
	return repositoryHandler.entityRepository.GetEntities(search, categories, limit, offset)
}

/*
GetEntity - Get an Entity from database and returns
a models. Entity on success
*/
func (e EntityRepository) GetEntity(id bson.ObjectId) (models.Entity, error) {
	return repositoryHandler.entityRepository.GetEntity(id)
}

/*
GetEntityBy - Get a Entity from database and returns
a models.Entity on success
*/
func (e EntityRepository) GetEntityBy(attribute string, value string) (models.Entity, error) {
	return repositoryHandler.entityRepository.GetEntityBy(attribute, value)
}

/*
GetEntityByPreviousTitle - get entity by previous title value
*/
func (e EntityRepository) GetEntityByPreviousTitle(title string, searchDate time.Time) (models.Entity, error) {
	return repositoryHandler.entityRepository.GetEntityByPreviousTitle(title, searchDate)
}

/*
TerminateEntity - terminate entity's life span by adding a tag to the entity title
*/
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
		if entityIsCompatible, existingEntity := (managers.EntityManager{}.CheckEntityCompatibility(existingEntity, entity)); entityIsCompatible {
			existingEntity = existingEntity.RemoveAttribute("new_title")
			log.Println("entity exists. terminating:", existingEntity.GetTitle())
			return repositoryHandler.entityRepository.UpdateEntity(existingEntity)
		}
	}
	return nil
}

/*
DeleteEntity - delete entity from the Database (This might break relations)
*/
func (e EntityRepository) DeleteEntity(entity models.Entity) error {
	return repositoryHandler.entityRepository.DeleteEntity(entity)
}

func (e EntityRepository) UpdateEntity(entity models.Entity) error {
	return repositoryHandler.entityRepository.UpdateEntity(entity)
}

func (e EntityRepository) NormalizeEntityTitle(entityTitle string) (string, error) {
	/*
		search for the title in the current system.
			get the search results from titles database
			for each search result match the string matching percentage
			pick the title with the highest percentage. that's the title of the entity
		if an acceptable title is not found in the database, try with normalize utility
			for each search result match the string matching percentage
			pick the title with the highest percentage. that's the title of the entity
		if an acceptable title is not found still,
			create entity with the existing name, tag it with a category name to identify
			add title to normalized name database
	*/
	normalizedTitle, isNormalized, processedEntityTitle := entityTitle, false, libraries.ProcessNameString(entityTitle)

	// try from existing normalization database
	normalizedNames, normalizedNameErr := NormalizedNameRepository{}.GetNormalizedNames(entityTitle, 1)

	if normalizedNameErr == nil {
		isNormalized, normalizedTitle = functions.SearchNormalizationInCache(normalizedNames, processedEntityTitle)
	}
	/*
		find an existing entity with matching name
	*/
	if !isNormalized {
		normalizedNames, normalizedNameErr := EntityRepository{}.GetEntities(entityTitle, nil, 1, 0)

		if normalizedNameErr == nil {
			isNormalized, normalizedTitle = functions.SearchNormalizationInDatabase(normalizedNames, processedEntityTitle)
		}
	}

	//try the search API
	if !isNormalized {
		isNormalized, normalizedTitle = functions.SearchNormalizationInSearchAPI(entityTitle, processedEntityTitle)
	}

	if isNormalized {
		log.Println("entity name normalized:", entityTitle, "->", normalizedTitle)
		return normalizedTitle, nil
	}

	//try the location API
	//if !isNormalized {
	//	isNormalized, normalizedTitle = functions.SearchNormalizationInLocationSearchAPI(entityTitle)
	//}

	if isNormalized {
		log.Println("entity name normalized:", entityTitle, "->", normalizedTitle)
		return normalizedTitle, nil
	}

	return entityTitle, errors.New(error_messages.NormalizationFailedError + " unable to find a match")
}

/*
GetStats Get entity states from the DB
*/
func (e EntityRepository) GetStats() (models.EntityStats, error) {
	return repositoryHandler.entityRepository.GetStats()
}

func (e EntityRepository) updateExistingEntity(entity models.Entity, existingEntity models.Entity) (models.Entity, error) {
	if existingEntity.GetImageURL() == "" {
		existingEntity = existingEntity.SetImageURL(entity.GetImageURL())
	}
	if existingEntity.GetSource() == "" {
		existingEntity = existingEntity.SetSource(entity.GetSource())
	}
	if existingEntity.GetSourceSignature() == "" {
		existingEntity = existingEntity.SetSourceSignature(entity.GetSourceSignature())
	}

	log.Println("entity exists. updating", existingEntity.GetTitle())
	existingEntity = existingEntity.SetSnippet()
	return existingEntity, repositoryHandler.entityRepository.UpdateEntity(existingEntity)
}

func (e EntityRepository) getExistingEntity(entity models.Entity) (models.Entity, error) {
	if entity.GetSourceDate().IsZero() {
		return e.GetEntityBy("title", entity.GetTitle())
	}
	return e.GetEntityByPreviousTitle(entity.GetTitle(), entity.GetSourceDate())
}

func (e EntityRepository) normalizeEntity(entity models.Entity) (models.Entity, string) {
	entityTitle := entity.GetTitle()
	if (managers.EntityManager{}.IsFromVerifiedSource(entity)) {
		NormalizedNameRepository{}.AddTitleToNormalizationDatabase(entity.GetTitle(), entity.GetTitle())
		return entity, entityTitle
	}
	entityTitle, normalizationErr := EntityRepository{}.NormalizeEntityTitle(entity.GetTitle())
	if normalizationErr == nil {
		return entity, entityTitle
	}
	entity = entity.AddCategory("arbitrary-entities")
	log.Println(error_messages.NormalizationFailedError, normalizationErr)
	return entity, entityTitle
}
