package parsers

import (
	"GIG/scripts/parsers"
)

func (t *TestParsers) TestThatPdfParserWorks() {
	result := parsers.ParsePdf("scripts/data/ahq_1005.pdf")
	t.AssertEqual(len(result), 88606)
}