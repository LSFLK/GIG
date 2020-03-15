package repositories

import (
	"GIG/app/models"
	"GIG/app/utilities/normalizers"
	"GIG/commons"
	"fmt"
	"strings"
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

func pickNormalizedNames(sourceName string, matchedString string, normalizedName string) (string, string) {
	if commons.StringsMatch(sourceName, matchedString, normalizers.StringMinMatchPercentage) {

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
				nameBeforeNormalizing, entity.Title = pickNormalizedNames(entity.GetTitle(),normalizedName.GetSearchText(), normalizedName.GetNormalizedText())
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
					nameBeforeNormalizing, entity.Title = pickNormalizedNames(entity.GetTitle(),normalizedName.GetTitle(), normalizedName.GetTitle())
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

				nameBeforeNormalizing, entity.Title = pickNormalizedNames(entity.GetTitle(),normalizedName, normalizedName)
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
