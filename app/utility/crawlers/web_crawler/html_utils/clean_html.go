package html_utils

import (
	"GIG/app/models"
	"GIG/app/utility"
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
		if !utility.StringInSlice(ignoreElements, n.Data) {
			endTag := ""
			trimmedData := strings.TrimSpace(n.Data)
			if n.Type == html.TextNode && trimmedData != "" {
				if !utility.StringInSlice(ignoreStrings, trimmedData) {
					result = result + n.Data
				}
			} else if n.Type == html.ElementNode {
				startTag := ""
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
					fixedURL := utility.FixUrl(href.Val, uri)
					if utility.ExtractDomain(uri) == "en.wikipedia.org" &&
						len(href.Val) > 0 &&
						string(href.Val[0]) != "#" &&
						title.Val != "" &&
						!utility.StringContainsAnyInSlice(ignoreTitles, title.Val) {

						linkedEntities = append(linkedEntities, models.Entity{Title: title.Val, SourceURL: fixedURL})

					}
					startTag = n.Data + " href='" + fixedURL + "' title='" + title.Val + "'"
				}
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

					fixedSrc := utility.FixUrl(src.Val, uri)
					fileName := utility.ExtractFileName(fixedSrc)
					bucketName := utility.ExtractDomain(fixedSrc)
					startTag = n.Data + " src='images/" + bucketName + "/" + fileName + "' width='" + width.Val + "'" + "' height='" + height.Val + "'"
					//startTag = n.Data + " src='" + fixedSrc + "' width='" + width.Val + "' height='" + height.Val + "'"
					imageList = append(imageList, models.Upload{Title: bucketName, SourceURL: fixedSrc})
				}
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
			if utility.StringInSlice(lineBreakers, n.Data) {
				result = result + "\n"
			}
		}
	}
	f(body)

	return result, linkedEntities, imageList
}
