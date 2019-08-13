// https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"GIG/app/models"
	"GIG/app/utility"
	"GIG/app/utility/entity_handlers"
	"GIG/app/utility/parsers"
	"GIG/app/utility/request_handlers"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/JackDanger/collectlinks"
	"net/url"
	"os"
	"strings"
)

/**
config before running
 */
var downloadDir = "app/cache/pdf_crawler/"
var standfordNERServer = "http://18.221.69.238:8080/classify"
var normalizeServer = "http://localhost:9000/api/normalize"
//var category = "Gazettes"
var category = "Tenders"

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("starting url not specified")
		os.Exit(1)
	}
	uri := args[0]

	resp, err := request_handlers.GetRequest(uri)

	if err != nil {
		panic(err)
	}

	links := collectlinks.All(strings.NewReader(resp))

	err = utility.EnsureDirectory(downloadDir)
	if err != nil {
		panic(err)
	}

	baseDir := downloadDir + utility.ExtractDomain(uri) + "/"
	for _, link := range links {
		if utility.FileTypeCheck(link, "pdf") {
			fmt.Println(link, uri)
			absoluteUrl := utility.FixUrl(link, uri)
			fmt.Println(absoluteUrl)

			// make directory if not exist
			if _, err := os.Stat(baseDir); os.IsNotExist(err) {
				if err != nil {
					fmt.Println(err)
				}
				createError := os.Mkdir(baseDir, os.ModePerm)
				if createError != nil {
					fmt.Println(createError)
				}
			}

			// download file
			encodedFileName := utility.ExtractFileName(absoluteUrl)
			filePath := baseDir + encodedFileName
			err := utility.DownloadFile(filePath, absoluteUrl)
			if err != nil {
				fmt.Println(err)
			}

			//parse pdf
			fileName, _ := url.QueryUnescape(encodedFileName)
			textContent := parsers.ParsePdf(filePath)
			fmt.Println(fileName)

			//NER extraction
			apiResp, apiErr := request_handlers.PostRequest(standfordNERServer, textContent)

			if apiErr != nil {
				fmt.Println(apiErr.Error())
			}
			var (
				entityTitles [][]string
				entities     []models.Entity
			)
			json.Unmarshal([]byte(apiResp), &entityTitles)

			//decode to entity
			entity := models.Entity{
				Title: utility.ExtractDomain(uri) + " - " + fileName,
			}.SetAttribute("", models.Value{
				Type:     "string",
				RawValue: textContent,
			}).AddCategory(category)

			for _, entityObject := range entityTitles {
				//normalize entity title before appending
				normalizedName, err := request_handlers.GetRequest(normalizeServer + "?searchText=" + entityObject[0])
				var response map[string]string
				json.Unmarshal([]byte(normalizedName), &response)
				if err == nil && response["status"] == "200" {
					entities = append(entities, models.Entity{Title: response["content"]}.AddCategory(entityObject[1]))
				}
			}

			entity, err = entity_handlers.AddEntitiesAsLinks(entity, entities)
			//save to db
			entity, saveErr := entity_handlers.CreateEntity(entity)
			if saveErr != nil {
				fmt.Println(saveErr.Error(), absoluteUrl)
			}
		}
	}

	fmt.Println("pdf crawling completed")

}
