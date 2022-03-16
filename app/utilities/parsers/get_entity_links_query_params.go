package parsers

import (
	"GIG-SDK/libraries"
	"github.com/revel/revel"
	"strconv"
)

func GetEntityLinksQueryParams(params *revel.Params) (error, int, int, []string) {
	limit, limitErr := strconv.Atoi(params.Values.Get("limit"))
	page, pageErr := strconv.Atoi(params.Values.Get("page"))
	attributes := params.Values.Get("attributes")
	attributesArray := libraries.ParseCategoriesString(attributes)

	if pageErr != nil || page < 1 {
		page = 1
	}

	return limitErr, page, limit, attributesArray
}
