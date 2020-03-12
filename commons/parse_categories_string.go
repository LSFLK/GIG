package commons

import "strings"

func ParseCategoriesString(categoriesString string)[]string{
	var categoriesArray []string

	if strings.TrimSpace(categoriesString) != "" {
		categoriesArray = strings.Split(categoriesString, ",")
	}
	return categoriesArray
}
