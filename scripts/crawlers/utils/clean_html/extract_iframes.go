package clean_html

import (
	"GIG/commons"
	"golang.org/x/net/html"
)

func ExtractIFrames(startTag string, n *html.Node, uri string) string {
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

