package create_entity

import (
	"GIG/app/models"
	"GIG/scripts/crawlers/utils"
	"GIG/scripts/entity_handlers"
	"time"
)

func CreateEntityFromText(textContent string, title string, categories []string, entityTitles []utils.NERResult) error {
	//decode to entity
	var entities []models.Entity
	entity := models.Entity{}.
		SetTitle(models.Value{Type: "string", RawValue: title, Source: "gazette", Date: time.Now()}).
		SetAttribute("", models.Value{
			Type:     "string",
			RawValue: textContent,
		}).AddCategories(categories)

	for _, entityObject := range entityTitles {
		//normalizedName, err := utils.NormalizeName(entityObject.EntityName)
		//if err == nil {
		entities = append(entities, models.Entity{Title: entityObject.EntityName}.AddCategory(entityObject.Category))
		//}
	}

	entity, err := entity_handlers.AddEntitiesAsLinks(entity, entities)
	if err != nil {
		panic(err)
	}

	//save to db
	entity, saveErr := entity_handlers.CreateEntity(entity)

	return saveErr
}
