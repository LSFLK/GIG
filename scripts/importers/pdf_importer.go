package main

import (
	"GIG/scripts/crawlers/pdf_crawler/create_entity"
	"GIG/scripts/crawlers/pdf_crawler/parsers"
	"GIG/scripts/crawlers/utils"
	"flag"
	"fmt"
	"os"
)

/**
config before running
 */

var pdfCategories = []string{"Gazette"}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("file path not specified")
		os.Exit(1)
	}
	filePath := args[0]
	//parse pdf
	textContent := parsers.ParsePdf(filePath)
	//NER extraction
	entityTitles, err := utils.ExtractEntityNames(textContent)
	if err != nil {
		fmt.Println(err)
	}
	if err := create_entity.CreateEntityFromText(textContent, "Gazette 2015", pdfCategories, entityTitles); err != nil {
		fmt.Println(err.Error(), filePath)
	}

	fmt.Println("pdf importing completed")

}
