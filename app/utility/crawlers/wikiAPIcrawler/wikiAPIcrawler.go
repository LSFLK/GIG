// https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"GIG/app/models"
	"GIG/app/utility/crawlers/wikiAPIcrawler/decoders"
	"GIG/app/utility/crawlers/wikiAPIcrawler/requests"
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
	contentResult, err := requests.GetContent(requests.PropTypeContent,title)
	if err != nil {
		return entity, err
	}
	linkResult, err := requests.GetContent(requests.PropTypeLinks,title)
	if err != nil {
		return entity, err
	}
	categoryResult, err := requests.GetContent(requests.PropTypeCategories,title)
	if err != nil {
		return entity, err
	}
	decoders.Decode(contentResult, &entity)
	decoders.Decode(linkResult, &entity)
	decoders.Decode(categoryResult, &entity)

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
