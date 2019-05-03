// https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"GIG/app/models"
	"GIG/app/utility"
	"GIG/app/utility/parsers"
	"GIG/app/utility/requesthandlers"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/JackDanger/collectlinks"
	"io/ioutil"
	"net/url"
	"os"
)

/**
config before running
 */
var apiUrl = "http://localhost:9000/api/add"
var downloadDir = "app/utility/crawlers/pdfcrawler/downloads/"
var standfordNERserver = "http://18.221.69.238:8080/classify"
var category = "Tenders"

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("starting url not specified")
		os.Exit(1)
	}
	uri := args[0]

	resp, err := requesthandlers.GetRequest(uri)

	if err != nil {
		panic(err)
	}

	links := collectlinks.All(resp.Body)

	baseDir := downloadDir + utility.ExtractDomain(uri) + "/"

	for _, link := range links {
		if utility.FileTypeCheck(link, "pdf") {
			absoluteUrl := utility.FixUrl(link, uri)

			// make directory if not exist
			if _, err := os.Stat(baseDir); os.IsNotExist(err) {
				os.Mkdir(baseDir, os.ModePerm)
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
			apiResp, apiErr := requesthandlers.PostRequest(standfordNERserver, textContent)
			defer apiResp.Body.Close()

			if apiErr != nil {
				fmt.Println(apiErr.Error())
			}
			body, readError := ioutil.ReadAll(apiResp.Body)
			if readError != nil {
				fmt.Println(readError.Error())
			}
			var entities [][]string
			json.Unmarshal(body, &entities)
			fmt.Println(entities)

			//decode to entity
			entity := models.Entity{}
			entity.Title = utility.ExtractDomain(uri) + " - " + fileName
			entity.SourceID = absoluteUrl
			entity.Content = textContent
			entity.Categories = append(entity.Categories, category) // change according to category crawling
			for _, classifiedClass := range entities {
				entity.Links = append(entity.Links, classifiedClass[0])
			}

			//save to db
			_, saveErr := requesthandlers.PostRequest(apiUrl, entity)
			if saveErr != nil {
				fmt.Println(saveErr.Error(), absoluteUrl)
			}

		}
	}

	fmt.Println("pdf crawling completed")

}
