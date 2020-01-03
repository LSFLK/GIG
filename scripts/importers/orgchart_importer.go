package main

import (
	"GIG/scripts/crawlers/pdf_crawler/create_entity"
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

var regexDate = regexp.MustCompile(`(0?[1-9]|[12][0-9]|3[01])*\.(0?[1-9]|1[012])\.\d{4}`)
var ministryTitle1 = regexp.MustCompile(`^\(\d+\)[ \t]+Minister of`)
var ministryTitle2 = regexp.MustCompile(`^Minister of`)
var departmentTitle = regexp.MustCompile(`^\d+\.`)
var departmentTitle2 = regexp.MustCompile(`\d+\.`)

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
	gazetteDate, err := ExtractGazetteDate(textContent)
	if err != nil {
		panic(err)
	}
	fmt.Println(gazetteDate)
	textContent= regexDate.ReplaceAllString(textContent,"")
	textContent = departmentTitle2.ReplaceAllStringFunc(textContent, func(m string) string {
		return departmentTitle2.ReplaceAllString(m, "\n"+m)
	})
	//fmt.Println(textContent)
	splitPage := strings.Split(textContent, parsers.NewPageMarker)

	dataStructure := make(map[string][]utils.NERResult)

	for _, page := range splitPage {
		splitArray := strings.Split(page, "\n")

		for index, line := range splitArray {
			line=StandardizeSpaces(line)
			ministryName, ministryFound := MinistryMatchFound(line)
			if ministryFound {

				i := 0
				j := len(splitArray)

				if index < 3*j/4 {
					i = index
				} else {
					j = index
				}

				startDepartments := false
				var departmentList []string
				for {
					if i == j {
						break
					}

					subLine := StandardizeSpaces(splitArray[i])
					nextMinistry, nextMinistryFound := MinistryMatchFound(subLine)
					if nextMinistryFound && nextMinistry != ministryName {
						break
					}

					if len(subLine) > 2 && (subLine[0:2] == "* " || subLine[0:2] == " (" || subLine[0:1] == "(") || (len(subLine) == 2 && strings.Contains(subLine,"x")) || (len(subLine) > 3 && startDepartments && subLine[0:2] == "1.") { // where department list ends
						startDepartments = false
					}

					if len(departmentTitle.FindStringSubmatch(subLine)) > 0 { // where department list is assumed to start
						startDepartments = true
					}

					if startDepartments {
						if len(departmentTitle.FindStringSubmatch(subLine)) > 0 { // identify numbered line
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
						dataStructure[ministryName] = append(dataStructure[ministryName], utils.NERResult{Category: "Department", EntityName: strings.TrimSpace(strings.Split(listLine, ".")[1])})
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

func ExtractGazetteDate(textContent string) (time.Time, error) {
	dateMatch := regexDate.FindStringSubmatch(textContent)
	if len(dateMatch) > 0 {
		return time.Parse("02.01.2006", dateMatch[0])
	}
	return time.Time{}, errors.New("unable to extract date of gazette")
}

func MinistryMatchFound(ministryName string) (string, bool) {
	ministryName = strings.TrimSpace(ministryName)
	ministryMatch1 := ministryTitle1.FindStringSubmatch(ministryName)
	ministryMatch2 := ministryTitle2.FindStringSubmatch(ministryName)
	if len(ministryMatch1) > 0 || len(ministryMatch2) > 0 {
		ministryNameArray := strings.Split(ministryName, ") ")
		if len(ministryNameArray) == 2 {
			ministryName = strings.TrimSpace(ministryNameArray[1])
		}
		return ministryName, true
	}
	return "", false
}

func StandardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

