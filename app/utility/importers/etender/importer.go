package main

import (
	"GIG/app/models"
	"GIG/app/utility/importers/etender/decoders"
	"GIG/app/utility/requesthandlers"
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
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
			companyEntity := models.Entity{
				Title: tender.Company,
			}.AddCategory("Organization")

			/**
			save company to db
			get company entity id
			set company id to tender.company attribute
			save tender entity
			 */
			var company models.Entity

			companyResponse, companySaveErr := requesthandlers.PostRequest(apiUrl, companyEntity)
			if companySaveErr != nil {
				fmt.Println(companySaveErr.Error(), companyEntity)
			}
			companyResponseBody, companyBodyError := ioutil.ReadAll(companyResponse.Body)
			if companyBodyError == nil {
				json.Unmarshal(companyResponseBody, &company)
			} else {
				fmt.Println(companySaveErr.Error(), companyEntity)
			}
			companyResponse.Body.Close()

			entity := decoders.MapToEntity(tender).AddCategory(category).
				SetAttribute("Company", models.Value{
					Type:     "objectId",
					RawValue: company.ID.Hex(),
				})
			resp, saveErr := requesthandlers.PostRequest(apiUrl, entity)
			if saveErr != nil {
				fmt.Println(saveErr.Error(), entity)
			}
			resp.Body.Close()
			fmt.Println(tender.Title, tender.Location)
		}
	}
}
