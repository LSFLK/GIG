package decoders

import (
	"GIG/app/models"
	"fmt"
)

func Decode(result map[string]interface{}, entity *models.Entity) {
	query := result["query"].(map[string]interface{})
	pages := query["pages"].(map[string]interface{})

	for _, page := range pages {

		pageObj := page.(map[string]interface{})

		if pageObj["extract"] != nil {
			fmt.Println("decoding content...")

			entity.Title = pageObj["title"].(string)
			entity.Content = pageObj["extract"].(string)
			entity.SourceID = fmt.Sprintf("wikiAPI%f", pageObj["pageid"])
		}

		if pageObj["links"] != nil {
			fmt.Println("decoding links...")
			links := pageObj["links"].([]interface{})

			for _, link := range links {
				linkObj := link.(map[string]interface{})
				entity.Links = append(entity.Links, linkObj["title"].(string))
			}
		}

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
