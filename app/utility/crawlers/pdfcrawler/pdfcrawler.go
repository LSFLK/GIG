// https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"GIG/app/models"
	"GIG/app/utility/parsers"
	"GIG/app/utility/requesthandlers"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/JackDanger/collectlinks"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var apiUrl = "http://18.221.69.238:9000/api/add"
var downloadDir = "app/utility/crawlers/pdfcrawler/downloads/"
var standfordNERserver = "http://18.221.69.238:8080/classify"

/**
	get page html and query body
	get notice link for pdf
	download pdf
	read pdf text/image
	extract pdf content using NER
	save to mongo
 */

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

	baseDir := downloadDir + getBaseDirectory(uri)

	for _, link := range links {
		if fileFormatOf(link, "pdf") {
			absoluteUrl := fixUrl(link, uri)

			// make directory if not exist
			if _, err := os.Stat(baseDir); os.IsNotExist(err) {
				os.Mkdir(baseDir, os.ModePerm)
			}

			// download file
			encodedFileName := getFileName(absoluteUrl)
			filePath := baseDir + encodedFileName
			err := DownloadFile(filePath, absoluteUrl)
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
			entity.Title = fileName
			entity.SourceID = absoluteUrl
			entity.Content = textContent
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

func fixUrl(href, base string) (string) {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseUrl.ResolveReference(uri)
	return uri.String()
}

func getFileName(link string) string {
	splitUrl := strings.Split(link, "/")
	return splitUrl[len(splitUrl)-1]
}

func getBaseDirectory(link string) string {
	splitUrl := strings.Split(link, "/")
	return splitUrl[2] + "/"
}

func fileFormatOf(link string, fileType string) bool {
	length := len(link)
	return length > 4 && link[length-len(fileType):length] == fileType
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
