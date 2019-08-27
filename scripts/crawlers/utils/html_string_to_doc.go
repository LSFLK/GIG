package utils

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func HTMLStringToDoc(resp string)(*goquery.Document, error){
	return goquery.NewDocumentFromReader(strings.NewReader(resp))
}
