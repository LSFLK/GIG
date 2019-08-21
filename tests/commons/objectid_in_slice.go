package commons

import (
	"GIG/commons"
	"gopkg.in/mgo.v2/bson"
)

func (t *TestCommons) TestThatObjectIdInSliceTestReturnsTrue() {
	testSlice := []bson.ObjectId{"1", "2", "3", "4"}
	t.AssertEqual(commons.ObjectIdInSlice(testSlice, "2"), true)
}

func (t *TestCommons) TestThatObjectIdInSliceTestReturnsFalse() {
	testSlice := []bson.ObjectId{"1", "2", "3", "4"}
	t.AssertEqual(commons.ObjectIdInSlice(testSlice, "5"), false)
}
