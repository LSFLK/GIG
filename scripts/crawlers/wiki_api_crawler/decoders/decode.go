package decoders

import (
	"GIG/app/models"
	"fmt"
	"strings"
	"time"
)

func Decode(result map[string]interface{}, entity *models.Entity) {
	query := result["query"].(map[string]interface{})
	pages := query["pages"].(map[string]interface{})

	for _, page := range pages {

		pageObj := page.(map[string]interface{})

		if pageObj["extract"] != nil {
			fmt.Println("	decoding content...")

			entity.Title = pageObj["title"].(string)
			tempEntity := entity.SetAttribute("", models.Value{
				Type:      "wikiText",
				RawValue:  pageObj["extract"].(string),
				Date: time.Now(),
			})
			entity.Attributes = tempEntity.Attributes
		}

		if pageObj["links"] != nil {
			fmt.Println("	decoding links...")
			links := pageObj["links"].([]interface{})

			for _, link := range links {
				linkObj := link.(map[string]interface{})
				entity.Links = append(entity.Links, linkObj["title"].(string))
			}
		}

		if pageObj["categories"] != nil {
			fmt.Println("	decoding categories...")
			categories := pageObj["categories"].([]interface{})

			for _, category := range categories {
				categoryObj := category.(map[string]interface{})
				categoryString := strings.Replace(categoryObj["title"].(string), "Category:", "", -1)
				entity.Categories = append(entity.Categories, categoryString)
			}
		}

	}

}
