package parsers

import (
	"GIG/scripts/crawlers/pdf_crawler/parsers"
)

func (t *TestParsers) TestThatPdfParserWorks() {
	result := parsers.ParsePdf("app/cache/test.pdf")
	t.AssertEqual(len(result), 124)
}