package decoders

import (
	"GIG/app/models"
)

func DecodeLinks(result map[string]interface{}, entity *models.Entity) {
	query := result["query"].(map[string]interface{})
	pages := query["pages"].(map[string]interface{})

	for _, page := range pages {

		pageObj := page.(map[string]interface{})
		links := pageObj["links"].(map[string]interface{})

		for _, link := range links{
			linkObj := link.(map[string]interface{})
			entity.Links= append(entity.Links, linkObj["title"].(string))
		}
	}

}
