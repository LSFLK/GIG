package clean_html

import (
	"GIG/app/models"
	"GIG/commons"
	"golang.org/x/net/html"
	"strings"
)

var defaultIgnoreElements = []string{"noscript", "script", "style", "input", "form","br","hr"}

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

func (c HtmlCleaner) CleanHTML(uri string, body *html.Node) (string, []models.Entity, []models.Upload, string) {
	var (
		result             string
		linkedEntities     []models.Entity
		f                  func(*html.Node)
		imageList          []models.Upload
		defaultImageSource string
		defaultImageWidth  int
	)

	lineBreakers := c.Config.LineBreakers
	ignoreElements := append(c.Config.IgnoreElements, defaultIgnoreElements...)
	ignoreStrings := c.Config.IgnoreStrings
	ignoreClasses := c.Config.IgnoreClasses

	f = func(n *html.Node) {
		if !commons.StringInSlice(ignoreElements, n.Data) && // ignore elements
			!commons.StringContainsAnyInSlice(ignoreClasses, ExtractClass(n)) { // ignore classes

			endTag := ""
			trimmedData := strings.TrimSpace(n.Data)
			if n.Type == html.TextNode && trimmedData != "" {
				if !commons.StringInSlice(ignoreStrings, trimmedData) {
					result = result + n.Data
				}
			} else if n.Type == html.ElementNode {
				startTag := ""
				imageSource := ""
				imageWidth := 0
				startTag, linkedEntities = c.extractLinks(startTag, n, uri, linkedEntities)
				startTag, imageList, imageSource, imageWidth = ExtractImages(startTag, n, uri, imageList)

				//set default image
				if imageSource != "" && (imageWidth > defaultImageWidth || defaultImageWidth == 0) {
					defaultImageWidth = imageWidth
					defaultImageSource = imageSource
				}

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

	return result, linkedEntities, imageList, defaultImageSource
}
