package commons

import (
	"GIG/commons"
)

func (t *TestCommons) TestThatStringInSliceTestReturnsTrue() {
	testSlice := []string{"1", "2", "3", "4"}
	t.AssertEqual(commons.StringInSlice(testSlice, "2"), true)
}

func (t *TestCommons) TestThatStringInSliceTestReturnsFalse() {
	testSlice := []string{"1", "2", "3", "4"}
	t.AssertEqual(commons.StringInSlice(testSlice, "5"), false)
}
