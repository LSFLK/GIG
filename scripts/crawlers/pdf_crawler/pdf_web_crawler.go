package main

import (
	"GIG/commons"
	"GIG/commons/request_handlers"
	"GIG/scripts/crawlers/pdf_crawler/create_entity"
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
var categories = []string{"Gazette"}

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

			fileName, _ := url.QueryUnescape(encodedFileName)
			_, err = create_entity.CreateEntityFromPdf(filePath, commons.ExtractDomain(uri)+" - "+fileName, categories)
			if err != nil {
				fmt.Println(err.Error(), absoluteUrl)
			}

		}
	}

	fmt.Println("pdf crawling completed")

}
