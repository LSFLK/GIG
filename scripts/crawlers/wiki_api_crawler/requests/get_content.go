package requests

import (
	"GIG/commons/request_handlers"
	"encoding/json"
	"net/url"
)

const apiURL = "https://en.wikipedia.org/w/api.php?action=query&format=json"

func GetContent(propType string, title string) (map[string]interface{}, error) {
	resp, err := request_handlers.GetRequest(generateURL(propType, title))
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(resp), &result)
	return result, err
}

func generateURL(propType string, title string) string {
	return apiURL + propType + "&titles=" + url.QueryEscape(title)
}
