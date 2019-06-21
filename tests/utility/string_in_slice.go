package utility

import (
	"GIG/app/utility"
)

func (t *TestUtilities) TestThatStringInSliceTestReturnsTrue() {
	testSlice := []string{"1", "2", "3", "4"}
	t.AssertEqual(utility.StringInSlice(testSlice, "2"), true)
}

func (t *TestUtilities) TestThatStringInSliceTestReturnsFalse() {
	testSlice := []string{"1", "2", "3", "4"}
	t.AssertEqual(utility.StringInSlice(testSlice, "5"), false)
}
