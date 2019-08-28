package clean_html

import (
	"GIG/app/models"
	"GIG/commons"
	"golang.org/x/net/html"
)

func ExtractImages(startTag string, n *html.Node, uri string, imageList []models.Upload) (string, []models.Upload) {
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
