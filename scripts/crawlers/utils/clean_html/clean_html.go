package clean_html

import (
	"GIG/app/models"
	"GIG/commons"
	"golang.org/x/net/html"
	"strings"
)

type Config struct {
	LineBreakers   []string
	IgnoreElements []string
	IgnoreStrings  []string
	IgnoreTitles   []string
	IgnoreClasses  []string
}

type HtmlCleaner struct {
	Config Config
}

func (c HtmlCleaner) CleanHTML(uri string, body *html.Node) (string, []models.Entity, []models.Upload) {
	var (
		result         string
		linkedEntities []models.Entity
		f              func(*html.Node)
		imageList      []models.Upload
	)

	lineBreakers := c.Config.LineBreakers
	ignoreElements := c.Config.IgnoreElements
	ignoreStrings := c.Config.IgnoreStrings
	ignoreClasses := c.Config.IgnoreClasses

	f = func(n *html.Node) {
		if !commons.StringInSlice(ignoreElements, n.Data) {

			//ignore if contains class
			if !commons.StringContainsAnyInSlice(ignoreClasses, ExtractClass(n)) {

				endTag := ""
				trimmedData := strings.TrimSpace(n.Data)
				if n.Type == html.TextNode && trimmedData != "" {
					if !commons.StringInSlice(ignoreStrings, trimmedData) {
						result = result + n.Data
					}
				} else if n.Type == html.ElementNode {
					startTag := ""
					startTag, linkedEntities = c.extractLinks(startTag, n, uri, linkedEntities)
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
	}
	f(body)

	return result, linkedEntities, imageList
}
