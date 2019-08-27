package utils

import (
	"GIG/commons/request_handlers"
	"GIG/scripts"
	"encoding/json"
	"github.com/revel/revel"
)

/**
normalize entity title before appending
 */
func NormalizeName(title string) (string, error) {

	normalizedName, err := request_handlers.GetRequest(scripts.NormalizeServer + "?searchText=" + title)
	if err != nil {
		return "", err
	}
	var response map[string]string
	if json.Unmarshal([]byte(normalizedName), &response); err != nil {
		return "", err
	}
	if response["status"] != "200" {
		return "", revel.NewErrorFromPanic("Server responded with" + response["status"])
	}
	return normalizedName, nil
}
