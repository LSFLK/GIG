package clean_html

import (
	"GIG/app/models"
	"GIG/commons"
	"golang.org/x/net/html"
	"strconv"
)

func ExtractImages(startTag string, n *html.Node, uri string, imageList []models.Upload) (string, []models.Upload, string, int) {
	sourceLink := ""
	imageWidth := 0
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
		sourceLink = "images/" + bucketName + "/" + fileName
		imageWidth, _ = strconv.Atoi(width.Val)
		startTag = n.Data + " src='" + sourceLink + "' width='" + width.Val + "'" + "' height='" + height.Val + "'"
		imageList = append(imageList, models.Upload{Title: bucketName, SourceURL: fixedSrc})
	}
	return startTag, imageList, sourceLink, imageWidth
}
