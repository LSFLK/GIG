package main

import (
	"GIG/app/models"
	"GIG/app/utility/requesthandlers"
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var apiUrl = "http://localhost:9000/api/add"
var category = "Tenders"

type ETender struct {
	Title       string    `json:"title"`
	Company     string    `json:"company"`
	SourceDate  time.Time `json:"source_date"`
	Category    string    `json:"category"`
	Subcategory string    `json:"subcategory"`
	Location    string    `json:"subcategory"`
	ClosingDate time.Time `json:"closing_date"`
	SourceName  string    `json:"source_name"`
	Description string    `json:"description"`
	Value       string    `json:"value"`
}

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
			sourceDate, _ := time.Parse("01/02/06", line[2])
			closingDate, _ := time.Parse("01/02/06", line[6])
			tender := ETender{
				Title:       line[0],
				Company:     line[1],
				SourceDate:  sourceDate,
				Category:    line[3],
				Subcategory: line[4],
				Location:    line[5],
				ClosingDate: closingDate,
				SourceName:  line[7],
				Description: line[8],
				Value:       line[9],
			}
			tenderJson, _ := json.Marshal(tender)
			entity := models.Entity{
				Title:    tender.Title +" - "+tender.Location,
				SourceID: "etenders.lk" + tender.Title + " " + tender.Location,
				Content:  string(tenderJson),
			}
			entity.Categories = append(entity.Categories, category)
			entity.Categories = append(entity.Categories, tender.Category)
			entity.Categories = append(entity.Categories, tender.Subcategory)
			entity.Links = append(entity.Links, tender.Company)
			entity.Links = append(entity.Links, tender.Location)

			//save to db
			_, saveErr := requesthandlers.PostRequest(apiUrl, entity)
			if saveErr != nil {
				fmt.Println(saveErr.Error(), entity)
			}
			fmt.Println(tender.Title, tender.Location)
		}
	}
}
