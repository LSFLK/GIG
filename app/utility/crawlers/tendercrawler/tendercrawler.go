// https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"GIG/app/utility/requesthandlers"
	"GIG/app/utility/parsers"
	"flag"
	"fmt"
	"github.com/JackDanger/collectlinks"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

//var apiUrl = "http://localhost:9000/api/add"
var downloadDir = "app/utility/crawlers/tendercrawler/downloads/"

/**
	get page html and query body
	get notice link for pdf
	download pdf
	read pdf text/image
	extract pdf content using NER/Tesseract
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

	baseDir:=downloadDir + getBaseDirectory(uri)
	for _, link := range links {
		if fileFormatOf(link, "pdf") {
			absolute := fixUrl(link, uri)

			// make directory is not exist
			if _, err := os.Stat(baseDir); os.IsNotExist(err) {
				os.Mkdir(baseDir, os.ModePerm)
			}

			// download file
			filePath := baseDir + getFileName(absolute)
			err := DownloadFile(filePath, absolute)
			if err != nil {
				fmt.Println(err)
			}

			//parse pdf
			fmt.Println(parsers.ParsePdf(filePath))

		}
	}

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
