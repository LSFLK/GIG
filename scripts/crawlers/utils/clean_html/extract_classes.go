package clean_html

import (
	"golang.org/x/net/html"
)

/*
return class attribute values of a html element
 */
func ExtractClass(n *html.Node) string {
	for _, attr := range n.Attr {
		if attr.Key == "class" {
			return attr.Val
		}
	}
	return ""
}
