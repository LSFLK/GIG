package normalizers

import (
	"GIG/app/utility/config"
	"GIG/app/utility/request_handlers"
	"encoding/json"
	"net/url"
	"strings"
)

var (
	SearchApiUrl = config.GetConfig().SearchApiUrl
	SearchAppKey = config.GetConfig().SearchAppKey
	Cx           = config.GetConfig().Cx
)

type SearchItem struct {
	Title string `json:"title"`
}

type SearchResponse struct {
	Items []SearchItem `json:"items"`
}

/**
given a text phrase returns the most matching entity names available
 */
func NormalizeName(searchString string) ([]string, error) {
	var (
		resultMap SearchResponse
		names     []string
	)
	result, err := request_handlers.GetRequest(SearchApiUrl + "?" + "cx=" + Cx + "&q=" + url.QueryEscape(searchString) + "&key=" + SearchAppKey)
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(result), &resultMap)
	for _, item := range resultMap.Items {
		names = append(names, strings.Replace(item.Title, " - Wikipedia", "", 1))
	}

	return names, nil
}
