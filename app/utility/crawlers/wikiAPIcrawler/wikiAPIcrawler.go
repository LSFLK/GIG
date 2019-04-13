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

	// todo: same title enqueued multiple times - fix issue
	for title := range queue {
		entity := enqueue(title, queue)
		_, err := requesthandlers.PostRequest(apiUrl, entity)
		if err != nil {
			fmt.Println(err.Error(), title)
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
			}
			decoders.Decode(result,&entity)
		}(propType)
	}
	wg.Wait()

	relatedTitles := append(entity.Categories, entity.Links...)
	for _, link := range relatedTitles {
		if link != "" {
			if !visited[link] {
				go func() { queue <- link }()
			}
		}
	}
	return entity
}
