package main

import (
	"GIG/scripts/crawlers/pdf_crawler/create_entity"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

/**
config before running
 */

var category = []string{"Gazette"}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("file path not specified")
		os.Exit(1)
	}
	filePath := args[0]

	_, err := create_entity.CreateEntityFromPdf(filePath, "Gazette 2017", category)
	if err != nil {
		fmt.Println(err.Error(), filePath)
	}

	fmt.Println("pdf importing completed")

}
