package create_entity

import (
	"GIG/app/models"
	"GIG/scripts/crawlers/pdf_crawler/parsers"
	"GIG/scripts/crawlers/utils"
	"GIG/scripts/entity_handlers"
	"fmt"
)

func CreateEntityFromPdf(filePath string, title string, category string) error {
	//parse pdf
	textContent := parsers.ParsePdf(filePath)

	//NER extraction
	entityTitles, err := utils.ExtractEntityNames(textContent)
	if err != nil {
		fmt.Println(err)
	}

	//decode to entity
	var entities []models.Entity
	entity := models.Entity{
		Title: title,
	}.SetAttribute("", models.Value{
		Type:     "string",
		RawValue: textContent,
	}).AddCategory(category)

	for _, entityObject := range entityTitles {
		//normalizedName, err := utils.NormalizeName(entityObject.EntityName)
		if err == nil {
			entities = append(entities, models.Entity{Title: entityObject.EntityName}.AddCategory(entityObject.Category))
		}
	}

	entity, err = entity_handlers.AddEntitiesAsLinks(entity, entities)

	//save to db
	entity, saveErr := entity_handlers.CreateEntity(entity)

	return saveErr
}