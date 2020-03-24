package normalizers

import (
	"GIG/commons/request_handlers"
	"encoding/json"
	"net/url"
	"strings"
)

var (
	SearchApiUrl string
	SearchAppKey string
	Cx           string
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
	result, err := request_handlers.GetRequest(SearchApiUrl + "?" + "cx=" + Cx + "&q=" + url.QueryEscape(searchString+" sri lanka") + "&key=" + SearchAppKey)
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(result), &resultMap)
	for _, item := range resultMap.Items {
		names = append(names, strings.Replace(item.Title, " - Wikipedia", "", 1))
	}

	return names, nil
}
