package pagination

import (
	"GIG-SDK/models"
	"GIG/app/repositories"
	"log"
)

func GetPaginatedEntityLinks(links []models.Link, attributesArray []string, page int, limit int) (error, []models.SearchResult) {
	var responseArray []models.SearchResult
	offset := (page - 1) * limit
	upperLimit := offset + limit
	if len(links) > offset {
		for i, link := range links {
			if i >= offset && i < upperLimit {
				var (
					linkedEntity models.Entity
					err          error
				)
				if len(link.GetDates()) > 0 {
					linkedEntity, err = repositories.EntityRepository{}.GetEntityByPreviousTitle(link.GetTitle(), link.GetDates()[0])
				} else {
					linkedEntity, err = repositories.EntityRepository{}.GetEntityBy("title", link.GetTitle())
				}

				if err != nil {
					log.Println(link.GetTitle(), err)
					return err, responseArray
				} else {
					responseArray = append(responseArray, models.SearchResult{}.ResultFrom(linkedEntity, attributesArray))
				}
			}
		}
	}

	return nil, responseArray
}
