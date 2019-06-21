package utility

import (
	"GIG/app/utility"
)

func (t *TestUtilities) TestThatFileTypeCheckTrueWorks() {
	link := "http://www.buildings.gov.lk/index.php"
	result := utility.FileTypeCheck(link,"php")
	t.AssertEqual(result, true)
}

func (t *TestUtilities) TestThatFileTypeCheckFalseWorks() {
	link := "http://www.buildings.gov.lk/index.php"
	result := utility.FileTypeCheck(link,"pdf")
	t.AssertEqual(result, false)
}