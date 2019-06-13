package requests

import (
	"GIG/app/utility/requesthandlers"
	"net/url"
)

const apiURL = "https://en.wikipedia.org/w/api.php?action=query&format=json"

func GetContent(propType string, title string) (map[string]interface{}, error) {
	return requesthandlers.GetJSON(generateURL(propType, title))
}

func generateURL(propType string, title string) string {
	return apiURL + propType + "&titles=" + url.QueryEscape(title)
}
