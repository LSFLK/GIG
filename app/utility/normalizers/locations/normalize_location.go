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

/**
given a text phrase returns the most matching available location
 */
func NormalizeLocation(searchString string) (map[string]interface{}, error) {
	result, err := request_handlers.GetRequest(ApiUrl + "?" + params + "&input=" + url.QueryEscape(searchString) + "&key=" + AppKey)
	if err != nil {
		return nil, err
	}
	var resultMap map[string]interface{}
	json.Unmarshal([]byte(result), &resultMap)
	return resultMap, nil
}
