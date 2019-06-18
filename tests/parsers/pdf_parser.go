package parsers

import (
	"GIG/app/utility/parsers"
	"github.com/revel/revel/testing"
)

type PdfParserTest struct {
	testing.TestSuite
}

func (t *PdfParserTest) Before() {
	println("Set up")
}

func (t *PdfParserTest) TestThatPdfParserWorks() {
	result := parsers.ParsePdf("app/cache/ahq_1005.pdf")
	t.AssertEqual(len(result), 88557)
}

func (t *PdfParserTest) After() {
	println("Tear down")
}
