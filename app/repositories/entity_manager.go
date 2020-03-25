package repositories

import (
	"GIG/app/models"
	"GIG/app/models/ValueType"
	"GIG/app/utilities/normalizers"
	"GIG/commons"
	"fmt"
	"github.com/pkg/errors"
)

func isFromVerifiedSource(entity models.Entity) bool {
	return entity.GetSourceSignature() == "trusted"
}

func NewEntityTitleIsWithinLifetimeOfExistingEntity(newTitleAttribute models.Attribute, lastTitleAttribute models.Attribute, entityIsTerminated bool) bool {
	return !(entityIsTerminated || newTitleAttribute.GetValue().GetDate().IsZero()) &&
		newTitleAttribute.GetValue().GetDate().After(lastTitleAttribute.GetValue().GetDate())
}
func NewEntityIsWithinLifeTimeOfExistingEntity(entity models.Entity, lastTitleAttribute models.Attribute, entityIsTerminated bool) bool {
	return !entityIsTerminated ||
		(entityIsTerminated && !entity.GetSourceDate().IsZero() &&
			entity.GetSourceDate().Before(lastTitleAttribute.GetValue().GetDate())) &&
			!entity.GetSourceDate().Before(lastTitleAttribute.GetValues()[0].GetDate())
}

func CheckEntityCompatibility(existingEntity models.Entity, entity models.Entity) (bool, models.Entity) {
	//if an entity exists
	if existingEntity.GetTitle() != "" {
		//if the entity has a "new_title" attribute use it to change the entity title
		newTitleAttribute, err := entity.GetAttribute("new_title")
		lastTitleAttribute, _ := existingEntity.GetAttribute("titles")

		isValidTitle := NewEntityTitleIsWithinLifetimeOfExistingEntity(newTitleAttribute, lastTitleAttribute, existingEntity.IsTerminated())

		isValidEntity := NewEntityIsWithinLifeTimeOfExistingEntity(entity, lastTitleAttribute, existingEntity.IsTerminated())

		if err == nil && isValidTitle {
			//add new title only if the new title date is before the date entity is terminated, else give an error
			fmt.Println("entity title modification found.", existingEntity.GetTitle(), "->", newTitleAttribute.GetValue().GetValueString())
			existingEntity = existingEntity.SetTitle(newTitleAttribute.GetValue())
		} else if err == nil && !isValidTitle {
			fmt.Println("new title cannot be assigned to a date after termination of the entity.")
		}

		if isValidEntity {

			if existingEntity.GetSourceDate().IsZero() && isFromVerifiedSource(entity) {
				existingEntity = existingEntity.SetSourceDate(entity.GetSourceDate()).
					SetTitle(models.Value{}.
						SetValueString(entity.GetTitle()).
						SetSource(entity.Source).
						SetDate(entity.GetSourceDate()).
						SetType(ValueType.String)).RemoveCategories([]string{"arbitrary-entities"})
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
			return true, existingEntity
		}
	}
	return false, existingEntity
}

func NormalizeEntityTitle(entityTitle string) (string, error) {
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
	normalizedTitle, isNormalized, processedEntityTitle := entityTitle, false, normalizers.ProcessNameString(entityTitle)

	// try from existing normalization database
	normalizedNames, normalizedNameErr := repositoryHandler.normalizedNameRepository.GetNormalizedNames(entityTitle, 1)

	if normalizedNameErr == nil {
		for _, normalizedName := range normalizedNames {
			if commons.StringsMatch(processedEntityTitle, normalizedName.GetSearchText(), normalizers.StringMinMatchPercentage) {
				isNormalized, normalizedTitle = true, normalizedName.GetNormalizedText()
				if isNormalized {
					fmt.Println("normalization found in cache", entityTitle, "->", normalizedTitle)
					break
				}
			}
		}
	}
	/**
	find an existing entity with matching name
	 */
	if !isNormalized {
		normalizedNames, normalizedNameErr := repositoryHandler.entityRepository.GetEntities(entityTitle, nil, 1, 0)

		if normalizedNameErr == nil {
			for _, normalizedName := range normalizedNames {
				if commons.StringsMatch(processedEntityTitle, normalizers.ProcessNameString(normalizedName.GetTitle()), normalizers.StringMinMatchPercentage) {
					isNormalized, normalizedTitle = true, normalizedName.GetTitle()
					if isNormalized {
						fmt.Println("normalization found in entity database", entityTitle, "->", normalizedTitle)
						break
					}
				}
			}
		}
	}

	//try the Wikipedia search API
	if !isNormalized {
		normalizedName, normalizedNameErr := normalizers.Normalize(entityTitle)
		if normalizedNameErr == nil && commons.StringsMatch(processedEntityTitle, normalizers.ProcessNameString(normalizedName), normalizers.StringMinMatchPercentage) {
			isNormalized, normalizedTitle = true, normalizedName
			fmt.Println("normalization found in search API", entityTitle, "->", normalizedTitle)
			AddTitleToNormalizationDatabase(entityTitle, normalizedTitle)
		} else {
			fmt.Println("normalization err:", normalizedNameErr)
		}
	}
	if isNormalized {
		fmt.Println("entity name normalized:", entityTitle, "->", normalizedTitle)
		return normalizedTitle, nil
	}

	return entityTitle, errors.New("normalization failed. unable to find a match")
}

func AddTitleToNormalizationDatabase(entityTitle string, normalizedName string) {
	// perform save in async
	go func(entityTitle string, normalizedName string) {
		_, err := NormalizedNameRepository{}.AddNormalizedName(
			models.NormalizedName{}.SetSearchText(entityTitle).SetNormalizedText(normalizedName),
		)
		if err != nil {
			fmt.Println("error while saving normalized title:", err)
		}
	}(entityTitle, normalizedName)
}
