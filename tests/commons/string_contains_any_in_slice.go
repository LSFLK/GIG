package commons

import (
	"GIG/commons"
)

func (t *TestCommons) TestThatStringContainsAnyInSliceTestReturnsTrue() {
	testSlice := []string{"1", "2", "3", "4","here"}
	t.AssertEqual(commons.StringContainsAnyInSlice(testSlice, "some value here"), true)
}

func (t *TestCommons) TestThatStringContainsAnyInSliceTestReturnsFalse() {
	testSlice := []string{"1", "2", "3", "4","some value here"}
	t.AssertEqual(commons.StringInSlice(testSlice, "else"), false)
}
