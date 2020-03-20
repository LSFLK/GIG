package model

import (
	"GIG/app/models"
	"GIG/app/models/ValueType"
	"strings"
	"time"
)

var (
	source0              = "source0"
	source               = "source"
	source2              = "source2"
	source3              = "source3"
	valueString          = "~test /tit?le % "
	valueString2         = "~test /tit?le % 2"
	valueString3         = "~test /tit?le % 3"
	date, _              = time.Parse("2006-1-2", "2010-5-20")
	date2, _             = time.Parse("2006-1-2", "2010-5-22")
	date3, _             = time.Parse("2006-1-2", "2011-5-22")
	valueType            = ValueType.String
	formattedValueString = "2test -title"
	testAttributeKey     = "test_attribute"

	testValueObj0 = models.Value{}.
		SetSource(source0).
		SetValueString(valueString).
		SetType(valueType)

	testValueObj = models.Value{}.
		SetSource(source).
		SetValueString(valueString).
		SetDate(date).
		SetType(valueType)

	testValueObj2 = models.Value{}.
		SetSource(source2).
		SetValueString(valueString2).
		SetDate(date2).
		SetType(valueType)

	testValueObj3 = models.Value{}.
		SetSource(source3).
		SetValueString(valueString3).
		SetDate(date3).
		SetType(valueType)
)

/*
entity set title works
 */
func (t *TestModels) TestThatEntitySetTitleWorks() {

	testEntity := models.Entity{}.SetTitle(testValueObj)
	titleAttribute, err := testEntity.GetAttribute("titles")

	titleValue := titleAttribute.GetValue()

	t.AssertEqual(err, nil)
	t.AssertEqual(testEntity.GetTitle(), formattedValueString)
	t.AssertEqual(titleValue.GetValueString(), formattedValueString)
	t.AssertEqual(titleValue.GetType(), valueType)
	t.AssertEqual(titleValue.GetDate(), date)
	t.AssertEqual(titleValue.GetSource(), source)
	t.Assert(titleValue.GetUpdatedDate().After(date))
}

/*
set attribute work for new attribute
 */
func (t *TestModels) TestThatEntitySetAttributeWorksForNewAttribute() {

	testEntity := models.Entity{}.SetAttribute(testAttributeKey, testValueObj)
	testAttribute, err := testEntity.GetAttribute(testAttributeKey)

	testValue := testAttribute.GetValue()

	t.AssertEqual(err, nil)
	t.AssertEqual(testValue.GetValueString(), strings.TrimSpace(valueString))
	t.AssertEqual(testValue.GetType(), valueType)
	t.AssertEqual(testValue.GetDate(), date)
	t.AssertEqual(testValue.GetSource(), source)
	t.Assert(testValue.GetUpdatedDate().After(date))
}

/*
set attribute works for existing attribute with same value
 */
func (t *TestModels) TestThatEntitySetAttributeWorksForExistingAttributeWithSameValue() {

	testEntity := models.Entity{}.SetAttribute(testAttributeKey, testValueObj)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj)

	testAttribute, err := testEntity.GetAttribute(testAttributeKey)
	testValue := testAttribute.GetValue()

	t.AssertEqual(err, nil)
	t.AssertEqual(len(testAttribute.GetValues()), 1)
	t.AssertEqual(testValue.GetValueString(), strings.TrimSpace(valueString))
	t.AssertEqual(testValue.GetType(), valueType)
	t.AssertEqual(testValue.GetDate(), date)
	t.AssertEqual(testValue.GetSource(), source)
	t.Assert(testValue.GetUpdatedDate().After(date))
}

/*
set attribute works for existing attribute with new value after the latest existing value
 */
func (t *TestModels) TestThatEntitySetAttributeWorksForExistingAttributeWithNewValueAfterLatestExistingValue() {

	testEntity := models.Entity{}.SetAttribute(testAttributeKey, testValueObj)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj2)

	testAttribute, err := testEntity.GetAttribute(testAttributeKey)
	testValue := testAttribute.GetValue()

	t.AssertEqual(err, nil)
	t.AssertEqual(len(testAttribute.GetValues()), 2)
	t.AssertEqual(testValue.GetValueString(), strings.TrimSpace(valueString2))
	t.AssertEqual(testValue.GetType(), valueType)
	t.AssertEqual(testValue.GetDate(), date2)
	t.AssertEqual(testValue.GetSource(), source2)
	t.Assert(testValue.GetUpdatedDate().After(date))
}

/*
set attribute works for existing attribute with new value in between the values
 */
func (t *TestModels) TestThatEntitySetAttributeWorksForExistingAttributeWithNewValueInBetweenExistingValues() {

	testEntity := models.Entity{}.SetAttribute(testAttributeKey, testValueObj)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj3)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj2)

	testAttribute, err := testEntity.GetAttribute(testAttributeKey)
	testValue := testAttribute.GetValues()[1]

	t.AssertEqual(err, nil)
	t.AssertEqual(len(testAttribute.GetValues()), 3)
	t.AssertEqual(testValue.GetValueString(), strings.TrimSpace(valueString2))
	t.AssertEqual(testValue.GetType(), valueType)
	t.AssertEqual(testValue.GetDate(), date2)
	t.AssertEqual(testValue.GetSource(), source2)
	t.Assert(testValue.GetUpdatedDate().After(date))
}

/*
set attribute works for existing attribute with new value before the first value date
 */
func (t *TestModels) TestThatEntitySetAttributeWorksForExistingAttributeWithNewValueBeforeTheFirstValue() {

	testEntity := models.Entity{}.SetAttribute(testAttributeKey, testValueObj3)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj2)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj)

	testAttribute, err := testEntity.GetAttribute(testAttributeKey)
	testValue := testAttribute.GetValues()[0]

	t.AssertEqual(err, nil)
	t.AssertEqual(len(testAttribute.GetValues()), 3)
	t.AssertEqual(testValue.GetValueString(), strings.TrimSpace(valueString))
	t.AssertEqual(testValue.GetType(), valueType)
	t.AssertEqual(testValue.GetDate(), date)
	t.AssertEqual(testValue.GetSource(), source)
	t.Assert(testValue.GetUpdatedDate().After(date))
}

/*
set attribute works for existing attribute with same value string but with zero date in existing value
 */
func (t *TestModels) TestThatEntitySetAttributeWorksForExistingAttributeWithSameValueButWithZeroDateInExistingValue() {

	testEntity := models.Entity{}.SetAttribute(testAttributeKey, testValueObj3)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj2)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj0)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj)

	testAttribute, err := testEntity.GetAttribute(testAttributeKey)
	testValue := testAttribute.GetValues()[0]

	t.AssertEqual(err, nil)
	t.AssertEqual(len(testAttribute.GetValues()), 3)
	t.AssertEqual(testValue.GetValueString(), strings.TrimSpace(valueString))
	t.AssertEqual(testValue.GetType(), valueType)
	t.AssertEqual(testValue.GetDate(), date)
	t.AssertEqual(testValue.GetSource(), source)
	t.Assert(testValue.GetUpdatedDate().After(date))
}

//TODO: write function to get attribute value by date

/*
TODO: cases ignored (unlikely to occur if data is accurate)
same value string past date
same value string future date
different values same date
 */
