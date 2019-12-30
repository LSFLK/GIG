package main

import (
	"GIG/scripts/crawlers/pdf_crawler/create_entity"
	"GIG/scripts/crawlers/pdf_crawler/parsers"
	"GIG/scripts/crawlers/utils"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

/**
config before running
 */

var category = []string{"OrgChart"}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("file path not specified")
		os.Exit(1)
	}
	filePath := args[0]
	textContent := parsers.ParsePdf(filePath) //parse pdf

	fmt.Println("processing org chart info...")
	splitPage := strings.Split(textContent, parsers.NewPageMarker)

	var regexDate = regexp.MustCompile(`(0?[1-9]|[12][0-9]|3[01])*\.(0?[1-9]|1[012])\.\d{4}`)
	var ministryTitle1 = regexp.MustCompile(`^\(\d+\) Minister of`)
	var ministryTitle2 = regexp.MustCompile(`^Minister of`)

	dateMatch := regexDate.FindStringSubmatch(textContent)
	if len(dateMatch) > 0 {
		gazetteDate, err := time.Parse("02.01.2006", dateMatch[0])
		if err != nil {
			panic(err)
		}
		fmt.Println(gazetteDate)
	} else {
		panic("Error: Unable to extract date of gazette")
	}

	dataStructure := make(map[string][]utils.NERResult)

	for _, page := range splitPage {
		splitArray := strings.Split(page, "\n")

		for _, line := range splitArray {
			ministryMatch1 := ministryTitle1.FindStringSubmatch(line)
			ministryMatch2 := ministryTitle2.FindStringSubmatch(line)
			if len(ministryMatch1) > 0 || len(ministryMatch2) > 0 {
				ministryName := strings.TrimSpace(line)
				ministryNameArray := strings.Split(line, ") ")
				if len(ministryNameArray) == 2 {
					ministryName = strings.TrimSpace(ministryNameArray[1])
				}
				i := 0
				startDepartments := false
				var departmentList []string
				for {
					if i == len(splitArray) {
						break
					}
					subline := splitArray[i]
					if len(subline) > 2 && (subline[0:2] == "* " || subline[0:2] == " (" || subline[0:1] == "(") { // where department list ends
						startDepartments = false
					}
					if startDepartments {
						if strings.Contains(subline, ". ") { // identify numbered line
							departmentList = append(departmentList, subline)
						} else {
							index := len(departmentList) - 1
							departmentList[index] = departmentList[index] + " " + subline
						}
					}
					if subline == "Corporations" && splitArray[i+1][0:1] != "(" && strings.Contains(splitArray[i+1], ". ") { // where department list is assumed to start
						startDepartments = true
					}
					//fmt.Println("		", subline)
					i++
				}

				if len(departmentList) > 0 {
					for _, listLine := range departmentList {
						dataStructure[ministryName] = append(dataStructure[ministryName], utils.NERResult{Category: "Department", EntityName: strings.Split(listLine, ". ")[1]})
					}
				}

			}
		}
	}

	for ministry, departments := range dataStructure {
		fmt.Println(ministry)
		for _, department := range departments {
			fmt.Println("	", department.EntityName)
		}
		if err := create_entity.CreateEntityFromText("", ministry, category, departments); err != nil {
			fmt.Println(err.Error(), filePath)
		}
	}

}
