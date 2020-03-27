package normalizers

import "encoding/xml"

type searchItem struct {
	Title string `json:"title"`
}

type searchResponse struct {
	Items []searchItem `json:"items"`
}


type xmlResponse struct {
	XMLName xml.Name `xml:"SearchSuggestion"`
	Section section `json:"Section"`
}

type section struct {
	XMLName xml.Name `xml:"Section"`
	Item []item `json:"Item"`
}

type item struct {
	XMLName xml.Name `xml:"Item"`
	Text string `json:"Text"`
}