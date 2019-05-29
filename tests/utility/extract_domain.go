package utility

import (
	"GIG/app/utility"
	"github.com/revel/revel/testing"
)

type ExtractDomainTest struct {
	testing.TestSuite
}

func (t *ExtractDomainTest) Before() {
	println("Set up")
}

func (t *ExtractDomainTest) TestThatExtractDomainWorks() {
	link := "http://www.buildings.gov.lk/index.php"
	result := utility.ExtractDomain(link)
	t.AssertEqual("www.buildings.gov.lk", result)
}

func (t *ExtractDomainTest) After() {
	println("Tear down")
}
