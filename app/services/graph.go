package services

import (
	models2 "GIG/app/models"
	"github.com/lsflk/gig-sdk/models"
)

func GetGraph(graph map[string]models2.GraphArray) (array map[string]map[string]int) {
	array = make(map[string]map[string]int)
	// find categories of all links and connect with categories of parent entity
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
	return
}
