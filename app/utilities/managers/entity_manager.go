package managers

import (
	"GIG-SDK/models"
	"GIG-SDK/models/ValueType"
	"fmt"
)

type EntityManager struct {
}

func (e EntityManager) IsFromVerifiedSource(entity models.Entity) bool {
	return entity.GetSourceSignature() == "trusted"
}

func (e EntityManager) NewEntityTitleIsWithinLifetimeOfExistingEntity(newTitleAttribute models.Attribute, lastTitleAttribute models.Attribute, entityIsTerminated bool) bool {
	return !(entityIsTerminated || newTitleAttribute.GetValue().GetDate().IsZero()) &&
		newTitleAttribute.GetValue().GetDate().After(lastTitleAttribute.GetValue().GetDate())
}
func (e EntityManager) NewEntityIsWithinLifeTimeOfExistingEntity(entity models.Entity, lastTitleAttribute models.Attribute, entityIsTerminated bool) bool {
	return !entityIsTerminated ||
		(entityIsTerminated && !entity.GetSourceDate().IsZero() &&
			entity.GetSourceDate().Before(lastTitleAttribute.GetValue().GetDate())) &&
			!entity.GetSourceDate().Before(lastTitleAttribute.GetValues()[0].GetDate())
}

func (e EntityManager) CheckEntityCompatibility(existingEntity models.Entity, entity models.Entity) (bool, models.Entity) {
	//if an entity exists
	if existingEntity.GetTitle() != "" {
		//if the entity has a "new_title" attribute use it to change the entity title
		newTitleAttribute, err := entity.GetAttribute("new_title")
		lastTitleAttribute, _ := existingEntity.GetAttribute("titles")

		isValidTitle := e.NewEntityTitleIsWithinLifetimeOfExistingEntity(newTitleAttribute, lastTitleAttribute, existingEntity.IsTerminated())

		isValidEntity := e.NewEntityIsWithinLifeTimeOfExistingEntity(entity, lastTitleAttribute, existingEntity.IsTerminated())

		if err == nil && isValidTitle {
			//add new title only if the new title date is before the date entity is terminated, else give an error
			fmt.Println("entity title modification found.", existingEntity.GetTitle(), "->", newTitleAttribute.GetValue().GetValueString())
			existingEntity = existingEntity.SetTitle(newTitleAttribute.GetValue())
		} else if err == nil && !isValidTitle {
			fmt.Println("new title cannot be assigned to a date after termination of the entity.")
		}

		if isValidEntity {

			if existingEntity.GetSourceDate().IsZero() && e.IsFromVerifiedSource(entity) {
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