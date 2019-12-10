package main

import (
	"GIG/scripts/crawlers/pdf_crawler/create_entity"
	"flag"
	"fmt"
	"os"
)

/**
config before running
 */
var category = "Gazette"

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("file path not specified")
		os.Exit(1)
	}
	filePath := args[0]

	if err := create_entity.CreateEntityFromPdf(filePath, "Gazette1", category); err != nil {
		fmt.Println(err.Error(), filePath)
	}

	fmt.Println("pdf importing completed")

}
