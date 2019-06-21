package utility

import (
	"GIG/app/utility"
)

func (t *TestUtilities) TestThatExtractDomainWorks() {
	link := "http://www.buildings.gov.lk/index.php"
	result := utility.ExtractDomain(link)
	t.AssertEqual("www.buildings.gov.lk", result)
}
