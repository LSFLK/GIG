// https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"GIG/app/models"
	"GIG/app/utility/crawlers/wiki_api_crawler/decoders"
	"GIG/app/utility/crawlers/wiki_api_crawler/requests"
	"GIG/app/utility/entity_handlers"
	"GIG/app/utility/request_handlers"
	"flag"
	"fmt"
	"os"
	"sync"
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

	lastTitle := ""
	for title := range queue {
		if title != lastTitle {
			lastTitle = title
			entity := enqueue(title, queue)
			resp, err := request_handlers.PostRequest(apiUrl, entity)
			resp.Body.Close()
			if err != nil {
				fmt.Println(err.Error(), title)
			}
		}
	}
}

func enqueue(title string, queue chan string) models.Entity {
	fmt.Println("fetching", title)
	visited[title] = true
	entity := models.Entity{}

	var wg sync.WaitGroup
	for _, propType := range requests.PropTypes() {

		wg.Add(1)
		go func(prop string) {
			defer wg.Done()
			result, err := requests.GetContent(prop, title)
			if err != nil {
				fmt.Println(err)
			} else {
				decoders.Decode(result, &entity)
			}
		}(propType)
	}
	wg.Wait()

	var tempEntity models.Entity
	for _, link := range entity.Links {
		if link != "" {
			if !visited[link] {
				go func() { queue <- link }()
			}
		}
		//add link as an entity
		linkEntity := models.Entity{Title: link}
		_, refVal, err := entity_handlers.AddEntityAsLink(tempEntity, linkEntity)

		if err != nil {
			fmt.Println("error creating link", link, entity)
		} else {
			tempEntity = tempEntity.AddLink(refVal)
		}

	}
	entity.Links = tempEntity.Links
	return entity
}
