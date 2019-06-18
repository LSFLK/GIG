package request_handlers

import (
	"encoding/json"
	"io/ioutil"
)

func GetJSON(uri string) (map[string]interface{}, error) {

	var result map[string]interface{}

	response, responseError := GetRequest(uri)
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
