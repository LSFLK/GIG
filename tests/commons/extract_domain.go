package commons

import (
	"GIG/commons"
)

func (t *TestCommons) TestThatExtractDomainWorks() {
	link := "http://www.buildings.gov.lk/index.php"
	result := commons.ExtractDomain(link)
	t.AssertEqual("www.buildings.gov.lk", result)
}
