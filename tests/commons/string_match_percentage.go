package commons

import "GIG/commons"

func (t *TestCommons) TestThatMatchStringWorksForEqualStrings() {
	matchPercent:=commons.StringMatchPercentage("long string","long string")
	t.AssertEqual(100, matchPercent)
}

func (t *TestCommons) TestThatMatchStringWorksForUnequalStrings() {
	matchPercent:=commons.StringMatchPercentage("long string","some other string")
	t.AssertEqual(17, matchPercent)
}

func (t *TestCommons) TestThatMatchStringWorksForPartiallyEqualStrings() {
	matchPercent:=commons.StringMatchPercentage("some what similar string","som wht similar stng")
	t.AssertEqual(66, matchPercent)
}
