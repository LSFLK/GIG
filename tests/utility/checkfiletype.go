package utility

import (
	"GIG/app/utility"
	"github.com/revel/revel/testing"
)

type CheckFileTypeTest struct {
	testing.TestSuite
}

func (t *CheckFileTypeTest) Before() {
	println("Set up")
}

func (t *CheckFileTypeTest) TestThatFileTypeCheckTrueWorks() {
	link := "http://www.buildings.gov.lk/index.php"
	result := utility.FileTypeCheck(link,"php")
	t.AssertEqual(result, true)
}

func (t *CheckFileTypeTest) TestThatFileTypeCheckFalseWorks() {
	link := "http://www.buildings.gov.lk/index.php"
	result := utility.FileTypeCheck(link,"pdf")
	t.AssertEqual(result, false)
}

func (t *CheckFileTypeTest) After() {
	println("Tear down")
}
