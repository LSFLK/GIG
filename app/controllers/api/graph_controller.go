package api

import (
	"GIG/app/constants/database"
	"GIG/app/controllers"
	"GIG/app/databases/mongodb_official"
	"github.com/lsflk/gig-sdk/models"
	"github.com/revel/revel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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
	graph := make(map[string]GraphArray)
	array := make(map[string]map[string]int)
	collection := mongodb_official.GetCollection(database.EntityCollection)
	findOptions := options.Find().SetProjection(bson.M{"title": 1, "links": 1, "categories": 1})
	cursor, err := collection.Find(mongodb_official.Context, bson.D{}, findOptions)
	if err != nil {
		return c.RenderJSON(controllers.BuildErrorResponse(err, 500))
	}
	// iterate through all documents and map to graph array
	for cursor.Next(mongodb_official.Context) {
		var entity models.Entity

		// Decode the document
		if err = cursor.Decode(&entity); err != nil {
			log.Println("cursor.Decode ERROR:", err)
			return c.RenderJSON(controllers.BuildErrorResponse(err, 500))
		}
		var links []string
		for _, link := range entity.Links {
			links = append(links, link.Title)
		}
		graph[entity.GetTitle()] = GraphArray{Title: entity.GetTitle(), Categories: entity.Categories, Links: links}
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
