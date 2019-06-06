package main

import (
	"GIG/app/models"
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
			entity := models.Entity{
				Title:    tender.Title + " - " + tender.Location,
				SourceID: "etenders.lk" + tender.Title + " " + tender.Location,
			}.
				AddCategory(category).
				AddCategory(tender.Category).
				AddCategory(tender.Subcategory).
				AddLink(tender.Company).
				AddLink(tender.Location).
				SetAttribute("Title", models.Value{
					Type:     "string",
					RawValue: tender.Title,
				}).
				SetAttribute("Company Name", models.Value{
					Type:     "string",
					RawValue: tender.Company,
				}).
				SetAttribute("Source Date", models.Value{
					Type:     "date",
					RawValue: tender.SourceDate.String(),
				}).
				SetAttribute("Category", models.Value{
					Type:     "string",
					RawValue: tender.Category,
				}).
				SetAttribute("Subcategory", models.Value{
					Type:     "string",
					RawValue: tender.Subcategory,
				}).
				SetAttribute("Source Date", models.Value{
					Type:     "date",
					RawValue: tender.SourceDate.String(),
				}).
				SetAttribute("Source Date", models.Value{
					Type:     "date",
					RawValue: tender.SourceDate.String(),
				}).
				SetAttribute("Source Date", models.Value{
					Type:     "date",
					RawValue: tender.SourceDate.String(),
				}).
				SetAttribute("Source Date", models.Value{
					Type:     "date",
					RawValue: tender.SourceDate.String(),
				}).
				SetAttribute("Source Date", models.Value{
					Type:     "date",
					RawValue: tender.SourceDate.String(),
				})

			//save to db
			_, saveErr := requesthandlers.PostRequest(apiUrl, entity)
			if saveErr != nil {
				fmt.Println(saveErr.Error(), entity)
			}
			fmt.Println(tender.Title, tender.Location)
		}
	}
}
