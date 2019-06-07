package main

import (
	"GIG/app/utility/importers/etender/decoders"
	"GIG/app/utility/requesthandlers"
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var apiUrl = "http://localhost:9000/api/add"
var category = "Tenders"

func main() {

	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("file path not specified")
		os.Exit(1)
	}
	filePath := args[0]

	csvFile, _ := os.Open(filePath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	ignoreHeaders := true

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		if ignoreHeaders {
			ignoreHeaders = false
		} else {
			tender := decoders.Decode(line)
			entity := decoders.MapToEntity(tender).AddCategory(category)

			//save to db
			resp, saveErr := requesthandlers.PostRequest(apiUrl, entity)
			if saveErr != nil {
				fmt.Println(saveErr.Error(), entity)
			}
			resp.Body.Close()
			fmt.Println(tender.Title, tender.Location)
		}
	}
}
