package normalizers

import (
	"GIG/app/utility/config"
	"GIG/app/utility/request_handlers"
	"encoding/json"
	"net/url"
)

var (
	MapApiUrl = config.GetConfig("map.api.url")
	MapAppKey = config.GetConfig("map.app.key")
	params = "inputtype=textquery&fields=name"
)

type Location struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

type MapGeometry struct {
	Location Location `json:"location"`
}

type MapResult struct {
	FormattedName string `json:"formatted_address"`
	Geometry MapGeometry `json:"geometry"`
}

type MapResponse struct {
	Status  string     `json:"status"`
	Results []MapResult `json:"results"`
}

/**
given a text phrase returns the most matching available locations
 */
func NormalizeLocation(searchString string) (MapResponse, error) {
	var resultMap MapResponse
	result, err := request_handlers.GetRequest(MapApiUrl + "?" + params + "&input=" + url.QueryEscape(searchString) + "&key=" + MapAppKey)
	if err != nil {
		return resultMap, err
	}

	json.Unmarshal([]byte(result), &resultMap)
	return resultMap, nil
}
