package decoders

import (
	"GIG/app/models"
	"fmt"
)

func DecodeLinks(result map[string]interface{}, entity *models.Entity) {
	query := result["query"].(map[string]interface{})
	pages := query["pages"].(map[string]interface{})

	for _, page := range pages {

		pageObj := page.(map[string]interface{})
		if pageObj["links"] != nil {
			fmt.Println("decoding links...")
			links := pageObj["links"].([]interface{})
			for _, link := range links {
				linkObj := link.(map[string]interface{})
				entity.Links = append(entity.Links, linkObj["title"].(string))
			}
		}
	}

}
