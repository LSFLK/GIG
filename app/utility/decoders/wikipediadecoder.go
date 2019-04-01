package decoders

import (
	"GIG/app/models"
	"github.com/PuerkitoBio/goquery"
	"io"
)

type WikipediaDecoder struct {
	Decoder
}

func (d WikipediaDecoder) DecodeSource(resp io.Reader) models.Entity {
	doc, _ := goquery.NewDocumentFromReader(resp)
	entity := models.Entity{}
	entity.Title = doc.Find("#firstHeading").First().Text()
	entity.Content = doc.Find("#mw-content-text").First().Text()
	return entity
}
