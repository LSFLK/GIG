package main

import (
	"GIG/scripts/crawlers/pdf_crawler/create_entity"
	"GIG/scripts/crawlers/pdf_crawler/parsers"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

/**
config before running
 */

var category = []string{"Gazette", "OrgChart"}

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
	if err := create_entity.CreateEntityFromText(filePath, "Gazette 2017", category); err != nil {
		fmt.Println(err.Error(), filePath)
	}

	fmt.Println("pdf importing completed")

	fmt.Println("processing org chart info...")
	textContent = strings.Replace(textContent, "\n", " ", -1)
	textContent = strings.Replace(textContent, "(", "\n(", -1)
	textContent = strings.Replace(textContent, "Column I Column II Column III Duties & Functions Departments, Statutory Laws to be Implemented Institutions  &  Public Corporations ", "", -1)
	textContent = strings.Replace(textContent, "  I  fldgi ", "", -1)
	splitArray := strings.Split(textContent, "\n")

	var ministryTitle = regexp.MustCompile(`^\((.[0-9]*?)\) Minister of`)
	var department = regexp.MustCompile(`(.[0-9]*?)\. `)
	for _, line := range splitArray {
		ministryMatch := ministryTitle.FindStringSubmatch(line)
		//departmentMatch := department.FindStringSubmatch(line)
		if len(ministryMatch) > 0 {
			s := department.ReplaceAllString(line, "\n")
			splitContent := strings.Split(s, "\n")
			fmt.Println(splitContent[0])
			for i, org := range splitContent {
				if i!=0 {
					fmt.Println("	", org)
				}
			}

		}
	}

}
