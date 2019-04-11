// https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"GIG/app/models"
	"GIG/app/utility/crawlers/wikiAPIcrawler/utils"
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
		response, _ := enqueue(title, queue)
		fmt.Println(response)
		//entity := decoder.DecodeSource(response, title)
		//fmt.Println(entity.Content)
		//_, err := requesthandlers.PostRequest(api_url, entity)
		//if err != nil {
		//	fmt.Println(err.Error(),title)
		//}
	}
}

func enqueue(title string, queue chan string) (models.Entity, error) {
	fmt.Println("fetching", title)
	visited[title] = true

	entity, err := utils.GetContent(title)
	// get content
	// get links
	// get categories

	if err != nil {
		return models.Entity{}, err
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