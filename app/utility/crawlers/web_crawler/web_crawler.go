// https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"GIG/app/models"
	"GIG/app/utility"
	"GIG/app/utility/entity_handlers"
	"GIG/app/utility/request_handlers"
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"os"
	"strings"
)

var visited = make(map[string]bool)
var apiUrl = "http://localhost:9000/api/add"

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
	if len(args) < 1 {
		fmt.Println("starting url not specified")
		os.Exit(1)
	}
	//decoder := WikipediaDecoder{}
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

	entity := models.Entity{SourceURL: uri}

	resp, err := request_handlers.GetRequest(uri)
	if err != nil {
		return entity, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp))
	if err != nil {
		return entity, err
	}

	entity.Title = doc.Find("#firstHeading").First().Text()
	bodyString, err := doc.Find("#bodyContent").First().Html()
	if err != nil {
		return entity, err
	}
	body, err := html.Parse(strings.NewReader(bodyString))
	if err != nil {
		return entity, err
	}

	result := ""
	lineBreakers := []string{"div", "caption"}
	ignoreElements := []string{"noscript", "script", "style", "input"}
	ignoreStrings := []string{"[", "]", "edit", "Jump to search", "Jump to navigation"}

	var linkedEntities []models.Entity
	var f func(*html.Node)
	f = func(n *html.Node) {
		if !utility.StringInSlice(ignoreElements, n.Data) {
			endTag := ""
			trimmedData := strings.TrimSpace(n.Data)
			if n.Type == html.TextNode && trimmedData != "" {
				if !utility.StringInSlice(ignoreStrings, trimmedData) {
					result = result + n.Data
				}
			} else if n.Type == html.ElementNode {
				startTag := ""
				if n.Data == "a" {
					var (
						href  html.Attribute
						title html.Attribute
					)
					for _, attr := range n.Attr {
						if attr.Key == "href" {
							href = attr
						} else if attr.Key == "title" {
							title = attr
						}
					}
					fixedURL := utility.FixUrl(href.Val, uri)
					if utility.ExtractDomain(uri) == "en.wikipedia.org" &&
						len(href.Val) > 0 &&
						string(href.Val[0]) != "#" &&
						title.Val != "" &&
						!strings.Contains(title.Val, ":") {

						linkedEntities = append(linkedEntities, models.Entity{Title: title.Val, SourceURL: fixedURL})
						if !visited[fixedURL] {
							go func() {
								queue <- fixedURL
							}()
						}
					}
					startTag = n.Data + " href='" + fixedURL + "' title='" + title.Val + "'"
				}
				if n.Data == "img" {
					var (
						src    html.Attribute
						width  html.Attribute
						height html.Attribute
					)
					for _, attr := range n.Attr {
						if attr.Key == "src" {
							src = attr
						} else if attr.Key == "width" {
							width = attr
						} else if attr.Key == "height" {
							height = attr
						}
					}
					fixedSrc := utility.FixUrl(src.Val, uri)
					imageString, err := request_handlers.GetRequest(fixedSrc)
					if err == nil {
						imgBase64Str := base64.StdEncoding.EncodeToString([]byte(imageString))
						startTag = n.Data + " src='data:image/png;base64," + imgBase64Str + "' width='" + width.Val + "'" + "' height='" + height.Val + "'"
					} else {
						startTag = n.Data + " src='" + fixedSrc + "' width='" + width.Val + "'" + "' height='" + height.Val + "'"
					}
				}
				if startTag == "" {
					result = result + "<" + n.Data + ">"
				} else {
					result = result + "<" + startTag + ">"
				}
				endTag = "</" + n.Data + ">"
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}

			if endTag != "" {
				result = result + endTag
			}
			if utility.StringInSlice(lineBreakers, n.Data) {
				result = result + "\n"
			}
		}
	}
	f(body)
	entity, err = entity_handlers.AddEntitiesAsLinks(entity, linkedEntities)
	fmt.Println(entity)
	entity = entity.SetAttribute("", models.Value{
		Type:     "html",
		RawValue: result,
	}).AddCategory("Wikipedia")
	return entity, nil
}
