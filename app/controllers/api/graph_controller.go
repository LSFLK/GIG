package api

import (
	"GIG/app/controllers"
	"GIG/app/repositories"
	"github.com/lsflk/gig-sdk/models"
	"github.com/revel/revel"
)

type GraphController struct {
	*revel.Controller
}
type GraphArray struct {
	Title            string
	Links            []string
	Categories       []string
	SourceCategories []string
}

func (c GraphController) GetGraph() revel.Result {

	array := make(map[string]map[string]int)
	graph, err := repositories.EntityRepository{}.GetGraph()

	if err != nil {
		return c.RenderJSON(controllers.BuildErrorResponse(err, 500))
	}

	// find categories of all links
	for _, item := range graph {
		categoryEntity := models.Entity{}
		for _, link := range item.Links {
			categoryEntity.AddCategories(graph[link].Categories)
		}
		item.SourceCategories = categoryEntity.Categories
		graph[item.Title] = item
		for _, category1 := range item.Categories {
			for _, category2 := range item.SourceCategories {
				if array[category1] == nil {
					array[category1] = make(map[string]int)
				}
				if array[category2] == nil {
					array[category2] = make(map[string]int)
				}
				array[category1][category2] += 1
				array[category2][category1] += 1
			}
		}
	}
	//connect link categories with categories of parent entity

	return c.RenderJSON(controllers.BuildSuccessResponse(array, 200))
}
