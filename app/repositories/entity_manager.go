package repositories

import (
	"GIG/app/models"
	"GIG/app/utilities/normalizers"
	"GIG/commons"
	"fmt"
	"strings"
	"time"
)

func isFromVerifiedSource(entity models.Entity) bool {
	return entity.SourceSignature == "trusted"
}

func newEntityTitleIsWithinLifetimeOfExistingEntity(newTitleAttribute models.Attribute, lastTitleValue models.Attribute, entityIsTerminated bool) bool {
	return (!entityIsTerminated) || (entityIsTerminated &&
		!newTitleAttribute.GetValue().Date.IsZero() &&
		newTitleAttribute.GetValue().Date.Before(lastTitleValue.GetValue().Date))
}
func newEntityIsWithinLifeTimeOfExistingEntity(entity models.Entity, lastTitleValue models.Attribute, entityIsTerminated bool) bool {
	return (!entityIsTerminated) || (entityIsTerminated &&
		!entity.SourceDate.IsZero() &&
		entity.SourceDate.Before(lastTitleValue.GetValue().Date))
}

func checkEntityCompatibility(existingEntity models.Entity, entity models.Entity) (bool, models.Entity) {
	//if an entity exists
	if existingEntity.GetTitle() != "" {
		//if the entity has a "new_title" attribute use it to change the entity title
		newTitleAttribute, err := entity.GetAttribute("new_title")
		entityIsTerminated := strings.Contains(existingEntity.GetTitle(), " - Terminated on ")
		lastTitleValue, _ := existingEntity.GetAttribute("titles")

		isValidTitle := newEntityTitleIsWithinLifetimeOfExistingEntity(newTitleAttribute, lastTitleValue, entityIsTerminated)

		isValidEntity := newEntityIsWithinLifeTimeOfExistingEntity(entity, lastTitleValue, entityIsTerminated)

		if err == nil && isValidTitle {
			//add new title only if the new title date is before the date entity is terminated, else give an error
			fmt.Println("entity title modification found.", existingEntity.GetTitle(), "->", newTitleAttribute.GetValue().GetValueString())
			existingEntity = existingEntity.SetTitle(newTitleAttribute.GetValue())
		} else if err == nil && !isValidTitle {
			fmt.Println("new title cannot be assigned to a date after termination of the entity.")
		}

		if isValidEntity {

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

			return true, existingEntity
		}
	}
	return false, existingEntity
}

func findExistingEntity(entity models.Entity) models.Entity {
	/**
	get entities containing title, select the entity matching the source date
		for each value matching the title. get the most recent date that is older than source date
			iterate each entity
				iterate each titles value
					if the value is the most recent then set the corresponding entity
	 */
	var mostRecentDate time.Time
	entitiesWithMatchingTitleAndDate, _ := repositoryHandler.entityRepository.GetEntityByPreviousState(entity.GetTitle(), entity.GetSourceDate())

	existingEntity := models.Entity{}
	// select the matching entity the given source date
	for _, resultEntity := range entitiesWithMatchingTitleAndDate {
		if resultAttribute, err := resultEntity.GetAttribute("titles"); err == nil {
			for _, resultValue := range resultAttribute.GetValues() {
				/**
				if titles match, if the source date is newer than title set date, source date is newer than most recent date
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
	return existingEntity
}

func pickNormalizedNames(sourceName string, normalizedName string) (string, string) {
	//TODO: compare entity title with search text, then add normalized name
	if commons.StringsMatch(sourceName, normalizedName, normalizers.StringMinMatchPercentage) {

		//if the entity signature is not found in the normalized names database, save it
		NormalizedNameRepository{}.AddNormalizedName(
			models.NormalizedName{}.SetSearchText(sourceName).SetNormalizedText(normalizedName),
		)
		return sourceName, normalizedName

	}
	return "", sourceName
}

func normalizeEntityTitle(entity models.Entity) models.Entity {
	/**
	search for the title in the current system.
		get the search results from titles database
		for each search result match the string matching percentage
		pick the title with highest percentage. that's the title of the entity
	if an acceptable title is not found in the database, try with normalize utility
		for each search result match the string matching percentage
		pick the title with highest percentage. that's the title of the entity
	if an acceptable title is not found still,
		create entity with the existing name, tag it with a category name to identify
		add title to normalized name database
	 */
	//TODO: if a trusted entity name, add it to the normalized name database
	if !isFromVerifiedSource(entity) {
		nameBeforeNormalizing := ""
		// try from existing normalization database
		normalizedNames, normalizedNameErr := repositoryHandler.normalizedNameRepository.GetNormalizedNames(entity.GetTitle(), 1)

		if normalizedNameErr == nil {
			for _, normalizedName := range normalizedNames {
				nameBeforeNormalizing, entity.Title = pickNormalizedNames(entity.GetTitle(), normalizedName.NormalizedText)
				if nameBeforeNormalizing != "" {
					break
				}
			}
		}
		/**
		find an existing entity with matching name
		 */
		if nameBeforeNormalizing == "" {
			normalizedNames, normalizedNameErr := repositoryHandler.entityRepository.GetEntities(entity.GetTitle(), nil, 1)

			if normalizedNameErr == nil {
				for _, normalizedName := range normalizedNames {
					nameBeforeNormalizing, entity.Title = pickNormalizedNames(entity.GetTitle(), normalizedName.GetTitle())
					if nameBeforeNormalizing != "" {
						break
					}
				}
			}
		}

		//try the Wikipedia search API
		if nameBeforeNormalizing == "" {
			normalizedName, normalizedNameErr := normalizers.Normalize(entity.GetTitle())
			if normalizedNameErr == nil {

				nameBeforeNormalizing, entity.Title = pickNormalizedNames(entity.GetTitle(), normalizedName)
			} else {
				fmt.Println("normalization err:", normalizedNameErr)
			}
		}
		if nameBeforeNormalizing == "" {
			entity = entity.AddCategory("arbitrary-entities")
		} else if nameBeforeNormalizing != entity.GetTitle() {
			fmt.Println("entity name normalized:", nameBeforeNormalizing, "->", entity.GetTitle())
		}
	}

	return entity
}
