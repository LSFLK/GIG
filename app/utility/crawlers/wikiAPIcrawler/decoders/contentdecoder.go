package decoders

import (
	"GIG/app/models"
	"fmt"
)

func DecodeContent(result map[string]interface{},entity *models.Entity ) {
	query := result["query"].(map[string]interface{})
	pages := query["pages"].(map[string]interface{})

	for _, page := range pages {

		pageObj := page.(map[string]interface{})
		entity.Title=pageObj["title"].(string)
		entity.Content=pageObj["extract"].(string)
		entity.SourceID=fmt.Sprintf("wikiAPI%f", pageObj["pageid"])
	}

}
