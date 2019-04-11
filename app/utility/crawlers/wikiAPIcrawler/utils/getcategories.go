package utils

import (
	"GIG/app/utility/requesthandlers"
	"net/http"
	"net/url"
)

func GetCategories(title string) (*http.Response, error) {

	uri := "https://en.wikipedia.org/w/api.php?action=query&format=json&prop=categories&titles="+url.QueryEscape(title)
	resp, err := requesthandlers.GetRequest(uri)
	return resp, err

}


