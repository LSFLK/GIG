package main

import (
	"GIG/app/models"
	"GIG/commons"
	"GIG/commons/request_handlers"
	"GIG/scripts/crawlers/pdf_crawler/parsers"
	"GIG/scripts/crawlers/utils"
	"GIG/scripts/entity_handlers"
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
var downloadDir = "scripts/crawlers/cache/"
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

	if err = commons.EnsureDirectory(downloadDir); err != nil {
		panic(err)
	}

	baseDir := downloadDir + commons.ExtractDomain(uri) + "/"
	for _, link := range links {
		if commons.FileTypeCheck(link, "pdf") {
			fmt.Println(link, uri)
			absoluteUrl := commons.FixUrl(link, uri)
			fmt.Println(absoluteUrl)

			// make directory if not exist
			if err = commons.EnsureDirectory(baseDir); err != nil {
				fmt.Println(err)
			}

			// download file
			encodedFileName := commons.ExtractFileName(absoluteUrl)
			filePath := baseDir + encodedFileName
			err := commons.DownloadFile(filePath, absoluteUrl)
			if err != nil {
				fmt.Println(err)
			}

			//parse pdf
			fileName, _ := url.QueryUnescape(encodedFileName)
			textContent := parsers.ParsePdf(filePath)
			fmt.Println(fileName)

			//NER extraction
			entityTitles, err := utils.ExtractEntityNames(textContent)
			if err != nil {
				fmt.Println(err)
			}

			//decode to entity
			var entities []models.Entity
			entity := models.Entity{
				Title: commons.ExtractDomain(uri) + " - " + fileName,
			}.SetAttribute("", models.Value{
				Type:     "string",
				RawValue: textContent,
			}).AddCategory(category)

			for _, entityObject := range entityTitles {
				normalizedName, err := utils.NormalizeName(entityObject.EntityName)
				if err == nil {
					entities = append(entities, models.Entity{Title: normalizedName}.AddCategory(entityObject.Category))
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
