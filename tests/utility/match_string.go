package utility

import "GIG/app/utility"

func (t *TestUtilities) TestThatMatchStringWorksForEqualStrings() {
	matchPercent:=utility.StringMatchPercentage("long string","long string")
	t.AssertEqual(100, matchPercent)
}

func (t *TestUtilities) TestThatMatchStringWorksForUnequalStrings() {
	matchPercent:=utility.StringMatchPercentage("long string","some other string")
	t.AssertEqual(17, matchPercent)
}

func (t *TestUtilities) TestThatMatchStringWorksForPartiallyEqualStrings() {
	matchPercent:=utility.StringMatchPercentage("some what similar string","som wht similar stng")
	t.AssertEqual(66, matchPercent)
}
