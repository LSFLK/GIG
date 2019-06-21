package parsers

import (
	"GIG/app/utility/parsers"
)

func (t *TestParsers) TestThatPdfParserWorks() {
	result := parsers.ParsePdf("app/cache/ahq_1005.pdf")
	t.AssertEqual(len(result), 88557)
}