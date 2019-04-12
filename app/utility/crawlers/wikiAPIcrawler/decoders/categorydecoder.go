package decoders

import (
	"GIG/app/models"
)

func DecodeCategories(result map[string]interface{}, entity *models.Entity) {
	query := result["query"].(map[string]interface{})
	pages := query["pages"].(map[string]interface{})

	for _, page := range pages {

		pageObj := page.(map[string]interface{})
		categories := pageObj["categories"].(map[string]interface{})

		for _, category := range categories{
			categoryObj := category.(map[string]interface{})
			entity.Categories= append(entity.Categories, categoryObj["title"].(string))
		}
	}

}
