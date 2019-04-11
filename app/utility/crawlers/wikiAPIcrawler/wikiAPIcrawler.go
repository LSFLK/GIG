// https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"GIG/app/models"
	"GIG/app/utility/crawlers/wikiAPIcrawler/decoders"
	"GIG/app/utility/crawlers/wikiAPIcrawler/utils"
	"GIG/app/utility/requesthandlers"
	"flag"
	"fmt"
	"os"
)

var visited = make(map[string]bool)
var api_url = "http://localhost:9000/api/add"

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
	if len(args) < 1 {
		fmt.Println("starting title not specified")
		os.Exit(1)
	}
	//decoder := decoders.WikiAPIDecoder{}
	queue := make(chan string)
	go func() { queue <- args[0] }()

	for title := range queue {
		entity, _ := enqueue(title, queue)
		_, err := requesthandlers.PostRequest(api_url, entity)
		if err != nil {
			fmt.Println(err.Error(),title)
		}
	}
}

func enqueue(title string, queue chan string) (models.Entity, error) {
	fmt.Println("fetching", title)
	visited[title] = true
	entity := models.Entity{}

	contentResult, err := utils.GetContent(title)
	decoders.DecodeSource(contentResult, &entity)

	// get links
	// get categories

	if err != nil {
		return entity, err
	}

	//for _, link := range links {
	//	if title != "" {
	//		if !visited[link] {
	//			go func() { queue <- link }()
	//		}
	//	}
	//}
	return entity, err
}
