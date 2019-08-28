package clean_html

import (
	"GIG/app/models"
	"GIG/commons"
	"golang.org/x/net/html"
	"strings"
)

var (
	lineBreakers   = []string{"div", "caption"}
	ignoreElements = []string{"noscript", "script", "style", "input"}
	ignoreStrings  = []string{"[", "]", "edit", "Jump to search", "Jump to navigation"}
	ignoreTitles   = []string{"(page does not exist)", ":"}
)

func CleanHTML(uri string, body *html.Node) (string, []models.Entity, []models.Upload) {
	var (
		result         string
		linkedEntities []models.Entity
		f              func(*html.Node)
		imageList      []models.Upload
	)

	f = func(n *html.Node) {
		if !commons.StringInSlice(ignoreElements, n.Data) {
			endTag := ""
			trimmedData := strings.TrimSpace(n.Data)
			if n.Type == html.TextNode && trimmedData != "" {
				if !commons.StringInSlice(ignoreStrings, trimmedData) {
					result = result + n.Data
				}
			} else if n.Type == html.ElementNode {
				startTag := ""
				startTag, linkedEntities = ExtractLinks(startTag, n, uri, linkedEntities)
				startTag, imageList = ExtractImages(startTag, n, uri, imageList)
				startTag = ExtractIFrames(startTag, n, uri)

				if startTag == "" {
					result = result + "<" + n.Data + ">"
				} else {
					result = result + "<" + startTag + ">"
				}
				endTag = "</" + n.Data + ">"
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}

			if endTag != "" {
				result = result + endTag
			}
			if commons.StringInSlice(lineBreakers, n.Data) {
				result = result + "\n"
			}
		}
	}
	f(body)

	return result, linkedEntities, imageList
}