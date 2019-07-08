package utility

import (
	"GIG/app/utility"
)

func (t *TestUtilities) TestThatStringContainsAnyInSliceTestReturnsTrue() {
	testSlice := []string{"1", "2", "3", "4","here"}
	t.AssertEqual(utility.StringContainsAnyInSlice(testSlice, "some value here"), true)
}

func (t *TestUtilities) TestThatStringContainsAnyInSliceTestReturnsFalse() {
	testSlice := []string{"1", "2", "3", "4","some value here"}
	t.AssertEqual(utility.StringInSlice(testSlice, "else"), false)
}
