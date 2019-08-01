package names

import (
	"GIG/app/utility/config"
	"GIG/app/utility/request_handlers"
	"encoding/json"
	"net/url"
	"strings"
)

var (
	ApiUrl = config.GetConfig().GoogleApiUrl
	AppKey = config.GetConfig().SearchAppKey
	Cx     = config.GetConfig().Cx
)

type Item struct {
	Title string `json:"title"`
}

type Response struct {
	Items []Item `json:"items"`
}

/**
given a text phrase returns the most matching entity names available
 */
func NormalizeName(searchString string) ([]string, error) {
	var (
		resultMap Response
		names     []string
	)
	result, err := request_handlers.GetRequest(ApiUrl + "?" + "cx=" + Cx + "&q=" + url.QueryEscape(searchString) + "&key=" + AppKey)
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(result), &resultMap)
	for _, item := range resultMap.Items {
		names = append(names, strings.Replace(item.Title, " - Wikipedia", "", 1))
	}

	return names, nil
}
