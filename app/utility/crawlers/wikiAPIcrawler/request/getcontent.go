package request

import (
	"GIG/app/utility/requesthandlers"
	"encoding/json"
	"io/ioutil"
	"net/url"
)

func GetContent(title string) (map[string]interface{}, error) {

	uri := "https://en.wikipedia.org/w/api.php?action=query&format=json&prop=extracts&exintro&explaintext&&titles=" + url.QueryEscape(title)
	var result map[string]interface{}

	response, responseError := requesthandlers.GetRequest(uri)
	if responseError != nil {
		return result, responseError
	}

	body, bodyError := ioutil.ReadAll(response.Body)
	if bodyError != nil {
		return result, bodyError
	}
	defer response.Body.Close()

	json.Unmarshal(body, &result)
	return result, bodyError

}
