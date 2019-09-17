// https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"GIG/app/models"
	"GIG/commons/request_handlers"
	"GIG/scripts/crawlers/utils"
	"GIG/scripts/crawlers/utils/clean_html"
	"GIG/scripts/crawlers/wiki_web_crawler/parsers"
	"GIG/scripts/entity_handlers"
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"os"
)

var visited = make(map[string]bool)

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
	if len(args) < 1 {
		fmt.Println("starting url not specified")
		os.Exit(1)
	}
	queue := make(chan string)
	go func() { queue <- args[0] }()

	for uri := range queue {
		entity, err := enqueue(uri, queue)
		if err != nil {
			fmt.Println("enqueue error:", err.Error(), uri)
		}
		_, err = entity_handlers.CreateEntity(entity)
		fmt.Println("entity added", entity.Title)
		if err != nil {
			fmt.Println(err.Error(), uri)
		}
	}
}

func enqueue(uri string, queue chan string) (models.Entity, error) {
	fmt.Println("fetching", uri)
	visited[uri] = true

	var (
		entity models.Entity
		err    error
		body   *html.Node
	)

	entity = models.Entity{SourceURL: uri}

	resp, err := request_handlers.GetRequest(uri)
	if err != nil {
		return entity, err
	}

	doc, err := utils.HTMLStringToDoc(resp)
	if err != nil {
		return entity, err
	}

	entity.Title, body, err = parsers.ParseHTMLContent(doc)
	if err != nil {
		return entity, err
	}

	//clean html code by removing unwanted information
	htmlCleaner := clean_html.HtmlCleaner{Config: clean_html.Config{
		LineBreakers:   []string{"div", "caption"},
		IgnoreElements: []string{"noscript", "script", "style", "input"},
		IgnoreStrings:  []string{"[", "]", "edit", "Jump to search", "Jump to navigation"},
		IgnoreTitles:   []string{"(page does not exist)", ":"},
	}}
	result, linkedEntities, imageList, defaultImageSource := htmlCleaner.CleanHTML(uri, body)
	entity.ImageURL = defaultImageSource

	// queue new links for crawling
	for _, linkedEntity := range linkedEntities {
		if !visited[linkedEntity.SourceURL] {
			go func(url string) {
				queue <- url
			}(linkedEntity.SourceURL)
		}
	}

	for _, image := range imageList {
		go func(payload models.Upload) {
			entity_handlers.UploadImage(payload)
		}(image)
	}

	// save linkedEntities (create empty if not exist)
	entity, err = entity_handlers.AddEntitiesAsLinks(entity, linkedEntities)
	entity = entity.SetAttribute("", models.Value{
		Type:     "html",
		RawValue: result,
	}).AddCategory("Wikipedia")

	return entity, nil
}
