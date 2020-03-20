package model

import (
	"GIG/app/models"
	"strings"
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

/*
TODO: cases ignored (unlikely to occur if data is accurate)
same value string past date
same value string future date
different values same date
 */
