package decoders

import (
	"GIG/app/models"
	"github.com/PuerkitoBio/goquery"
	"io"
)

type WikipediaDecoder struct {
}

func (d WikipediaDecoder) DecodePage(resp io.Reader) models.Entity {
	doc, _ := goquery.NewDocumentFromReader(resp)
	entity := models.Entity{}
	entity.Title = doc.Find("title").First().Text()
	entity.Content = doc.Find("h1").First().Text()
	return entity
}
