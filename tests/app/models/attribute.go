package models

import (
	"GIG-SDK//models"
	"time"
)

/*
entity set title works
 */
func (t *TestModels) TestThatAttributeGetValueByDateWorks() {

	testEntity := models.Entity{}.SetAttribute(testAttributeKey, testValueObj3)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj2)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj0)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj)

	testAttribute, err := testEntity.GetAttribute(testAttributeKey)

	testDate, _ := time.Parse("2006-1-2", "2010-5-18")
	testDate1, _ := time.Parse("2006-1-2", "2010-5-21")
	testDate2, _ := time.Parse("2006-1-2", "2010-5-22")
	testDate3, _ := time.Parse("2006-1-2", "2010-8-30")
	testDate4, _ := time.Parse("2006-1-2", "2012-8-30")

	t.AssertEqual(err, nil)
	t.AssertEqual(testAttribute.GetValueByDate(testDate).GetSource(), "")
	t.AssertEqual(testAttribute.GetValueByDate(testDate1).GetSource(), source)
	t.AssertEqual(testAttribute.GetValueByDate(testDate2).GetSource(), source2)
	t.AssertEqual(testAttribute.GetValueByDate(testDate3).GetSource(), source2)
	t.AssertEqual(testAttribute.GetValueByDate(testDate4).GetSource(), source3)
}
