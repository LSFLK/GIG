package parsers

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"strings"
)

func ParseHTMLContent(doc *goquery.Document) (string, *html.Node, error) {
	title := doc.Find("#firstHeading").First().Text()
	bodyString, err := doc.Find("#bodyContent").First().Html()
	if err != nil {
		return title, nil, err
	}
	body, err := html.Parse(strings.NewReader(bodyString))
	if err != nil {
		return title, nil, err
	}

	return title, body, err
}
