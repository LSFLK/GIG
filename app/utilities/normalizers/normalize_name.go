package normalizers

import (
	"encoding/xml"
	"github.com/lsflk/gig-sdk"
	"net/url"
)

//var (
//	SearchApiUrl string
//	SearchAppKey string
//	Cx           string
//)

// Using Wikipedia Open Search
func NormalizeName(searchString string) ([]string, error) {
	var (
		response xmlResponse
		names    []string
	)
	urlString := "https://en.wikipedia.org/w/api.php?action=opensearch&limit=5&format=xml&search=" + url.QueryEscape(searchString)
	result, err := GIG_SDK.GigClient{}.GetRequest(urlString)
	if err != nil {
		return nil, err
	}

	xml.Unmarshal([]byte(result), &response)

	for _, suggestedName := range response.Section.Item {
		names = append(names, suggestedName.Text)
	}

	return names, nil
}

/**
given a text phrase returns the most matching entity names available
 */
// Using google custom search engine
//func NormalizeName(searchString string) ([]string, error) {
//	var (
//		resultMap SearchResponse
//		names     []string
//	)
//	result, err := request_handlers.GetRequest(SearchApiUrl + "?" + "cx=" + Cx + "&q=" + url.QueryEscape(searchString+" sri lanka") + "&key=" + SearchAppKey)
//	if err != nil {
//		return nil, err
//	}
//	json.Unmarshal([]byte(result), &resultMap)
//	if len(resultMap.Items) == 0 {
//		log.Println("url", SearchApiUrl+"?"+"cx="+Cx+"&q="+url.QueryEscape(searchString+" sri lanka")+"&key="+SearchAppKey)
//		return nil, errors.New("search API returned error message.")
//	}
//	for _, item := range resultMap.Items {
//		names = append(names, strings.Replace(item.Title, " - Wikipedia", "", 1))
//	}
//
//	return names, nil
//}
