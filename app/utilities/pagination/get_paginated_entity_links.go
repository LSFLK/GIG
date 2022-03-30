package pagination

import (
	"GIG/app/repositories"
	"github.com/lsflk/gig-sdk/models"
	"log"
)

func GetPaginatedEntityLinks(links []models.Link, attributesArray []string, page int, limit int) (error, []models.SearchResult) {
	var (
		responseArray []models.SearchResult
	)

	offset := (page - 1) * limit
	upperLimit := offset + limit

	if len(links) > offset {
		for i, link := range links {
			if i >= offset && i < upperLimit {
				_, responseArray = addLinkedEntityToResults(link, responseArray, attributesArray)
			}
		}
	}

	return nil, responseArray
}

func addLinkedEntityToResults(link models.Link, responseArray []models.SearchResult, attributesArray []string) (error, []models.SearchResult) {
	linkedEntity, err := getLinkedEntity(link)

	if err != nil {
		log.Println(link.GetTitle(), err)
		return err, responseArray
	}
	return nil, append(responseArray, models.SearchResult{}.ResultFrom(linkedEntity, attributesArray))
}

func getLinkedEntity(link models.Link) (models.Entity, error) {
	if len(link.GetDates()) > 0 {
		return repositories.EntityRepository{}.GetEntityByPreviousTitle(link.GetTitle(), link.GetDates()[0])
	}
	return repositories.EntityRepository{}.GetEntityBy("title", link.GetTitle())
}
