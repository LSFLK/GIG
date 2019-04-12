package decoders

import (
	"GIG/app/models"
	"fmt"
)

func DecodeCategories(result map[string]interface{}, entity *models.Entity) {
	query := result["query"].(map[string]interface{})
	pages := query["pages"].(map[string]interface{})

	for _, page := range pages {

		pageObj := page.(map[string]interface{})
		if pageObj["categories"] != nil {
			fmt.Println("decoding categories...")
			categories := pageObj["categories"].([]interface{})

			for _, category := range categories {
				categoryObj := category.(map[string]interface{})
				entity.Categories = append(entity.Categories, categoryObj["title"].(string))
			}
		}
	}

}
