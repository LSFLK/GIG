package clean_html

import (
	"GIG/app/models"
	"GIG/commons"
	"golang.org/x/net/html"
	"time"
)

func (c HtmlCleaner) extractLinks(startTag string, n *html.Node, uri string, linkedEntities []models.Entity) (string, []models.Entity) {
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
			!commons.StringContainsAnyInSlice(c.Config.IgnoreTitles, title.Val) {

			linkedEntities = append(linkedEntities,
				models.Entity{
					SourceURL: fixedURL,
				}.
					SetTitle(models.Value{
						Type:      "string",
						RawValue:  title.Val,
						Date: time.Now(),
					}))

		}
		startTag = n.Data + " href='" + fixedURL + "' title='" + title.Val + "'"
	}

	return startTag, linkedEntities
}
