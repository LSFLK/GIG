package utils

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
				startTag, linkedEntities = extractLinks(startTag, n, uri, linkedEntities)
				startTag, imageList = extractImages(startTag, n, uri, imageList)
				startTag = extractIFrames(startTag, n, uri)

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

func extractLinks(startTag string, n *html.Node, uri string, linkedEntities []models.Entity) (string, []models.Entity) {
	if n.Data == "a" {
		var (
			href  html.Attribute
			title html.Attribute
		)
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				href = attr
			} else if attr.Key == "title" {
				title = attr
			}
		}
		fixedURL := commons.FixUrl(href.Val, uri)
		if commons.ExtractDomain(uri) == "en.wikipedia.org" &&
			len(href.Val) > 0 &&
			string(href.Val[0]) != "#" &&
			title.Val != "" &&
			!commons.StringContainsAnyInSlice(ignoreTitles, title.Val) {

			linkedEntities = append(linkedEntities, models.Entity{Title: title.Val, SourceURL: fixedURL})

		}
		startTag = n.Data + " href='" + fixedURL + "' title='" + title.Val + "'"
	}

	return startTag, linkedEntities
}

func extractImages(startTag string, n *html.Node, uri string, imageList []models.Upload) (string, []models.Upload) {
	if n.Data == "img" {
		var (
			src    html.Attribute
			width  html.Attribute
			height html.Attribute
		)
		for _, attr := range n.Attr {
			if attr.Key == "src" {
				src = attr
			} else if attr.Key == "width" {
				width = attr
			} else if attr.Key == "height" {
				height = attr
			}
		}

		fixedSrc := commons.FixUrl(src.Val, uri)
		fileName := commons.ExtractFileName(fixedSrc)
		bucketName := commons.ExtractDomain(fixedSrc)
		startTag = n.Data + " src='images/" + bucketName + "/" + fileName + "' width='" + width.Val + "'" + "' height='" + height.Val + "'"
		imageList = append(imageList, models.Upload{Title: bucketName, SourceURL: fixedSrc})
	}
	return startTag, imageList
}

func extractIFrames(startTag string, n *html.Node, uri string) string {
	if n.Data == "iframe" {
		var (
			src    html.Attribute
			width  html.Attribute
			height html.Attribute
		)
		for _, attr := range n.Attr {
			if attr.Key == "src" {
				src = attr
			} else if attr.Key == "width" {
				width = attr
			} else if attr.Key == "height" {
				height = attr
			}
		}

		fixedSrc := commons.FixUrl(src.Val, uri)
		startTag = n.Data + " src='" + fixedSrc + "' width='" + width.Val + "'" + "' height='" + height.Val + "'"
	}
	return startTag
}
