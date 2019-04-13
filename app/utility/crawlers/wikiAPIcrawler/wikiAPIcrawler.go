// https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"GIG/app/models"
	"GIG/app/utility/crawlers/wikiAPIcrawler/decoders"
	"GIG/app/utility/requesthandlers"
	"flag"
	"fmt"
	"os"
)

var visited = make(map[string]bool)
var apiUrl = "http://localhost:9000/api/add"

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
	if len(args) < 1 {
		fmt.Println("starting title not specified")
		os.Exit(1)
	}
	queue := make(chan string)
	go func() { queue <- args[0] }()

	// todo: same title enqueued multiple times - fix issue
	for title := range queue {
		entity, _ := enqueue(title, queue)
		_, err := requesthandlers.PostRequest(apiUrl, entity)
		if err != nil {
			fmt.Println(err.Error(), title)
		}
	}
}

func enqueue(title string, queue chan string) (models.Entity, error) {
	fmt.Println("fetching", title)
	visited[title] = true
	entity := models.Entity{}

	// todo: async the 3 requests
	contentResult, err := GetContent(PropTypeContent,title)
	if err != nil {
		return entity, err
	}
	linkResult, err := GetContent(PropTypeLinks,title)
	if err != nil {
		return entity, err
	}
	categoryResult, err := GetContent(PropTypeCategories,title)
	if err != nil {
		return entity, err
	}
	decoders.DecodeContent(contentResult, &entity)
	decoders.DecodeLinks(linkResult, &entity)
	decoders.DecodeCategories(categoryResult, &entity)

	relatedTitles:=append(entity.Categories,entity.Links...)
	for _, link := range relatedTitles {
		if link != "" {
			if !visited[link] {
				go func() { queue <- link }()
			}
		}
	}
	return entity, err
}
