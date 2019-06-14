package utility

import (
	"GIG/app/utility"
	"github.com/revel/revel/testing"
)

type StringInSliceTest struct {
	testing.TestSuite
}

func (t *StringInSliceTest) Before() {
	println("Set up")
}

func (t *StringInSliceTest) TestThatStringInSliceTestReturnsTrue() {
	testSlice := []string{"1", "2", "3", "4"}
	t.AssertEqual(utility.StringInSlice(testSlice, "2"), true)
}

func (t *StringInSliceTest) TestThatStringInSliceTestReturnsFalse() {
	testSlice := []string{"1", "2", "3", "4"}
	t.AssertEqual(utility.StringInSlice(testSlice, "5"), false)
}

func (t *StringInSliceTest) After() {
	println("Tear down")
}
