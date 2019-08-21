package commons

import (
	"GIG/commons"
)

func (t *TestCommons) TestThatFileTypeCheckTrueWorks() {
	link := "http://www.buildings.gov.lk/index.php"
	result := commons.FileTypeCheck(link,"php")
	t.AssertEqual(result, true)
}

func (t *TestCommons) TestThatFileTypeCheckFalseWorks() {
	link := "http://www.buildings.gov.lk/index.php"
	result := commons.FileTypeCheck(link,"pdf")
	t.AssertEqual(result, false)
}