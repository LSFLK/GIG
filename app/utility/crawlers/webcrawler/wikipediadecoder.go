package main

import (
	"GIG/app/models"
	"github.com/PuerkitoBio/goquery"
	"io"
)

type WikipediaDecoder struct {
}

func (d WikipediaDecoder) DecodeSource(resp io.Reader, uri string) models.Entity {
	doc, _ := goquery.NewDocumentFromReader(resp)
	entity := models.Entity{}
	entity.SourceID = uri
	entity.Title = doc.Find("#firstHeading").First().Text()
	entity.Content = doc.Find("#mw-content-text").First().Text()
	return entity
}
