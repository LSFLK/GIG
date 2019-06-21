package utility

import (
	"GIG/app/utility"
	"gopkg.in/mgo.v2/bson"
)

func (t *TestUtilities) TestThatObjectIdInSliceTestReturnsTrue() {
	testSlice := []bson.ObjectId{"1", "2", "3", "4"}
	t.AssertEqual(utility.ObjectIdInSlice(testSlice, "2"), true)
}

func (t *TestUtilities) TestThatObjectIdInSliceTestReturnsFalse() {
	testSlice := []bson.ObjectId{"1", "2", "3", "4"}
	t.AssertEqual(utility.ObjectIdInSlice(testSlice, "5"), false)
}
