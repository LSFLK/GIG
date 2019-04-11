package decoders

import (
	"GIG/app/models"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
)

type WikiAPIDecoder struct {
	Decoder
}

func (d WikiAPIDecoder) DecodeSource(resp io.Reader, uri string) models.Entity {
	doc, _ := goquery.NewDocumentFromReader(resp)
	fmt.Println(doc)
	entity := models.Entity{}
	entity.Source = uri
	entity.Title = doc.Find("#firstHeading").First().Text()
	entity.Content = doc.Find("#mw-content-text").First().Text()
	return entity
}
