package normalizers

import (
	"encoding/json"
	"github.com/lsflk/gig-sdk/client"
	"log"
	"net/url"
)

var (
	MapApiUrl string
	MapAppKey string
	params    = "inputtype=textquery&fields=formatted_address"
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
	//Geometry MapGeometry `json:"geometry"`
}

type MapResponse struct {
	Status  string      `json:"status"`
	Results []MapResult `json:"candidates"`
}

/**
given a text phrase returns the most matching available locations
 */
func NormalizeLocation(searchString string) (MapResponse, error) {
	var resultMap MapResponse
	result, err := client.GigClient{}.GetRequest(MapApiUrl + "?" + params + "&input=" + url.QueryEscape(searchString) + "&key=" + MapAppKey)
	if err != nil {
		return resultMap, err
	}

	json.Unmarshal([]byte(result), &resultMap)
	log.Println(resultMap)
	return resultMap, nil
}
