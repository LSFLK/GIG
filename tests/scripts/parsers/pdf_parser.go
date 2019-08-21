package parsers

import (
	"GIG/app/scripts/parsers"
)

func (t *TestParsers) TestThatPdfParserWorks() {
	result := parsers.ParsePdf("app/scripts/data/ahq_1005.pdf")
	t.AssertEqual(len(result), 88606)
}