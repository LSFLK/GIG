package main

import (
	"GIG/scripts/crawlers/pdf_crawler/parsers"
	"GIG/scripts/crawlers/utils"
	"errors"
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

	var ministryTitle1 = regexp.MustCompile(`^\(\d+\)[ \t]+Minister of`)
	var ministryTitle2 = regexp.MustCompile(`^Minister of`)

	gazetteDate, err := ExtractGazetteDate(textContent)
	if err != nil {
		panic(err)
	}
	fmt.Println(gazetteDate)

	dataStructure := make(map[string][]utils.NERResult)

	for _, page := range splitPage {
		splitArray := strings.Split(page, "\n")
		//numberOfLines := len(splitArray)

		for _, line := range splitArray {
			ministryName := strings.TrimSpace(line)
			ministryMatch1 := ministryTitle1.FindStringSubmatch(ministryName)
			ministryMatch2 := ministryTitle2.FindStringSubmatch(ministryName)
			if len(ministryMatch1) > 0 || len(ministryMatch2) > 0 {
				ministryNameArray := strings.Split(ministryName, ") ")
				if len(ministryNameArray) == 2 {
					ministryName = strings.TrimSpace(ministryNameArray[1])
				}
				//if index>(3*numberOfLines/4){
				//	fmt.Println(ministryName, "bottom")
				//}else{
				//	fmt.Println(ministryName, "top")
				//}

				i := 0
				startDepartments := false
				var departmentList []string
				for {
					if i == len(splitArray) {
						break
					}
					subLine := splitArray[i]
					if len(subLine) > 2 && (subLine[0:2] == "* " || subLine[0:2] == " (" || subLine[0:1] == "(") || (len(subLine) > 3 && startDepartments && subLine[0:3] == "1. ") { // where department list ends
						if startDepartments{
							departmentList = append(departmentList, "1. ***")
						}
						startDepartments = false
					}

					if len(subLine) > 3 && subLine[0:3] == "1. " { // where department list is assumed to start
						startDepartments = true
					}

					if startDepartments {
						if strings.Contains(subLine, ". ") { // identify numbered line
							departmentList = append(departmentList, subLine)
						} else {
							index := len(departmentList) - 1
							departmentList[index] = departmentList[index] + " " + subLine
						}
					}
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
			fmt.Println("	", department.EntityName, len(department.EntityName))
		}
		//if err := create_entity.CreateEntityFromText("", ministry, category, departments); err != nil {
		//	fmt.Println(err.Error(), filePath)
		//}
	}

}

func ExtractGazetteDate(textContent string) (time.Time, error) {
	var regexDate = regexp.MustCompile(`(0?[1-9]|[12][0-9]|3[01])*\.(0?[1-9]|1[012])\.\d{4}`)
	dateMatch := regexDate.FindStringSubmatch(textContent)
	if len(dateMatch) > 0 {
		return time.Parse("02.01.2006", dateMatch[0])
	}
	return time.Time{}, errors.New("unable to extract date of gazette")
}

func ExtractDepartmentsFromGazette() {
	//i := 0
	//startDepartments := false
	//var departmentList []string
	//for {
	//	if i == len(splitArray)-1 {
	//		break
	//	}
	//	subline := splitArray[i]
	//	if len(subline) > 2 && (subline[0:2] == "* " || subline[0:2] == " (" || subline[0:1] == "(") { // where department list ends
	//		startDepartments = false
	//	}
	//	if startDepartments {
	//		if strings.Contains(subline, ". ") { // identify numbered line
	//			departmentList = append(departmentList, subline)
	//		} else {
	//			index := len(departmentList) - 1
	//			departmentList[index] = departmentList[index] + " " + subline
	//		}
	//	}
	//	if subline == "Corporations" && splitArray[i+1][0:1] != "(" && strings.Contains(splitArray[i+1], ". ") { // where department list is assumed to start
	//		startDepartments = true
	//	}
	//	i++
	//}
	//
	//if len(departmentList) > 0 {
	//	for _, listLine := range departmentList {
	//		dataStructure[ministryName] = append(dataStructure[ministryName], utils.NERResult{Category: "Department", EntityName: strings.Split(listLine, ". ")[1]})
	//	}
	//}
}
