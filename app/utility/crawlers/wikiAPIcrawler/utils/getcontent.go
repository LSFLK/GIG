package utils

import (
	"GIG/app/models"
	"GIG/app/utility/requesthandlers"
	"encoding/json"
	"fmt"
	"net/url"
)

type ContentResponse struct {
	BatchComplete string `json:"batchcomplete"`
	Query         Query  `json:"query"`
}

type Query struct {
	Pages NumObj `json:"pages"`
}
type NumObj struct {
	Obj Page `json:"26750"`
}

type Page struct {
	PageId  int    `json:"pageid"`
	NS      int    `json:"ns"`
	Title   string `json:"title"`
	Extract string `json:"extract"`
}

func GetContent(title string) (models.Entity, error) {

	uri := "https://en.wikipedia.org/w/api.php?action=query&format=json&prop=extracts&exintro&explaintext&&titles=" + url.QueryEscape(title)
	resp, err := requesthandlers.GetRequest(uri)
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	defer resp.Body.Close()
	testObj := ContentResponse{}
	json.NewDecoder(resp.Body).Decode(&testObj)
	fmt.Println(testObj)
	if err != nil {
		return models.Entity{}, err // todo: return null array
	}

	return models.Entity{}, err

}
