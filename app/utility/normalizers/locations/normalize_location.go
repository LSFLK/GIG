package locations

import (
	"GIG/app/utility/config"
	"GIG/app/utility/request_handlers"
	"encoding/json"
	"net/url"
)

var (
	ApiUrl = config.GetConfig().MapApiUrl
	AppKey = config.GetConfig().AppKey
	params = "inputtype=textquery&fields=name"
)

type Location struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

type Geometry struct {
	Location Location `json:"location"`
}

type Result struct {
	FormattedName string `json:"formatted_address"`
	Geometry Geometry `json:"geometry"`
}

type Response struct {
	Status  string     `json:"status"`
	Results []Result `json:"results"`
}

/**
given a text phrase returns the most matching available locations
 */
func NormalizeLocation(searchString string) (Response, error) {
	var resultMap Response
	result, err := request_handlers.GetRequest(ApiUrl + "?" + params + "&input=" + url.QueryEscape(searchString) + "&key=" + AppKey)
	if err != nil {
		return resultMap, err
	}

	json.Unmarshal([]byte(result), &resultMap)
	return resultMap, nil
}
