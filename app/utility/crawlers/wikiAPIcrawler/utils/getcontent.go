package utils

import (
	"GIG/app/models"
	"GIG/app/utility/requesthandlers"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	//fmt.Println(string(body))

	//testObj := ContentResponse{}
	//json.NewDecoder(resp.Body).Decode(&testObj)
	query := result["query"].(map[string]interface{})
	pages := query["pages"].(map[string]interface{})
	for _, value := range pages {
		// Each value is an interface{} type, that is type asserted as a string
		pageId := value.(map[string]interface{})
		fmt.Println(pageId["title"])
	}

	if err != nil {
		return models.Entity{}, err // todo: return null array
	}

	return models.Entity{}, err

}
