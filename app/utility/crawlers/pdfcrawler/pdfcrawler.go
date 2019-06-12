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
var apiUrl = "http://18.221.69.238:9000/api/add"
var downloadDir = "app/cache/pdfcrawler/"
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
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	links := collectlinks.All(resp.Body)

	// make directory if not exist
	if _, err := os.Stat(downloadDir); os.IsNotExist(err) {
		if err != nil {
			panic(err)
		}
		createError := os.Mkdir(downloadDir, os.ModePerm)
		if createError != nil {
			panic(createError)
		}
	}

	baseDir := downloadDir + utility.ExtractDomain(uri) + "/"
	fmt.Println("here", baseDir)
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
			apiResp, apiErr := requesthandlers.PostRequest(standfordNERserver, textContent)

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
			apiResp.Body.Close()
			//decode to entity
			entity := models.Entity{
				Title:    utility.ExtractDomain(uri) + " - " + fileName,
				SourceID: absoluteUrl,
			}.SetAttribute("", models.Value{
				Type:     "string",
				RawValue: textContent,
			}).AddCategory(category)
			for _, classifiedClass := range entities {
				entity = entity.AddLink(classifiedClass[0])
			}

			//save to db
			saveResp, saveErr := requesthandlers.PostRequest(apiUrl, entity)
			if saveErr != nil {
				fmt.Println(saveErr.Error(), absoluteUrl)
			}
			saveResp.Body.Close()

		}
	}

	fmt.Println("pdf crawling completed")

}
