package managers

import (
	"GIG-SDK/models"
	"GIG/app/utilities/managers"
	"GIG/tests/app/test_values"
)

/*
New entity title is within lifetime of existing entity returns false if entity is terminated
 */
func (t *TestManagers) TestThatNewEntityTitleIsWithinLifetimeOfExistingEntityReturnsFalseIfEntityIsTerminated() {
	TestValue := managers.EntityManager{}.NewEntityTitleIsWithinLifetimeOfExistingEntity(models.Attribute{}, models.Attribute{}, true)

	t.AssertEqual(TestValue, false)
}

/*
New entity title is within lifetime of existing entity returns true if new title date is after last title date
 */
func (t *TestManagers) TestThatNewEntityTitleIsWithinLifetimeOfExistingEntityReturnsTrueIfNewTitleDateIsAfterLastTitleDate() {

	lastTitleAttribute := models.Attribute{}.SetValue(test_values.TestValueObj)
	newTitleAttribute := models.Attribute{}.SetValue(test_values.TestValueObj2)

	TestValue := managers.EntityManager{}.NewEntityTitleIsWithinLifetimeOfExistingEntity(newTitleAttribute, lastTitleAttribute, false)

	t.AssertEqual(newTitleAttribute.GetValue().GetDate().After(lastTitleAttribute.GetValue().GetDate()), true)
	t.AssertEqual(TestValue, true)
}

/*
New entity title is within lifetime of existing entity returns false if new title date is before last title date
 */
func (t *TestManagers) TestThatNewEntityTitleIsWithinLifetimeOfExistingEntityReturnsFalseIfNewTitleDateIsBeforeLastTitleDate() {

	lastTitleAttribute := models.Attribute{}.SetValue(test_values.TestValueObj2)
	newTitleAttribute := models.Attribute{}.SetValue(test_values.TestValueObj)

	TestValue := managers.EntityManager{}.NewEntityTitleIsWithinLifetimeOfExistingEntity(newTitleAttribute, lastTitleAttribute, false)

	t.AssertEqual(newTitleAttribute.GetValue().GetDate().Before(lastTitleAttribute.GetValue().GetDate()), true)
	t.AssertEqual(TestValue, false)
}

/*
New entity title is within lifetime of existing entity returns false if new title date zero
 */
func (t *TestManagers) TestThatNewEntityTitleIsWithinLifetimeOfExistingEntityReturnsFalseIfNewTitleDateIsZero() {

	lastTitleAttribute := models.Attribute{}.SetValue(test_values.TestValueObj2)
	newTitleAttribute := models.Attribute{}.SetValue(test_values.TestValueObj0)

	TestValue := managers.EntityManager{}.NewEntityTitleIsWithinLifetimeOfExistingEntity(newTitleAttribute, lastTitleAttribute, false)

	t.AssertEqual(newTitleAttribute.GetValue().GetDate().IsZero(), true)
	t.AssertEqual(TestValue, false)
}

/*
New entity title is within lifetime of existing entity returns false if new title date equals last title date
 */
func (t *TestManagers) TestThatNewEntityTitleIsWithinLifetimeOfExistingEntityReturnsFalseIfNewTitleDateEqualsLastTitleDate() {

	lastTitleAttribute := models.Attribute{}.SetValue(test_values.TestValueObj2)
	newTitleAttribute := models.Attribute{}.SetValue(test_values.TestValueObj2)

	TestValue := managers.EntityManager{}.NewEntityTitleIsWithinLifetimeOfExistingEntity(newTitleAttribute, lastTitleAttribute, false)

	t.AssertEqual(newTitleAttribute.GetValue().GetDate(), lastTitleAttribute.GetValue().GetDate())
	t.AssertEqual(TestValue, false)
}

/*
New entity is within lifetime of existing entity returns
true if existing entity is not terminated and new entity source date is after existing entity source date
 */
func (t *TestManagers) TestThatNewEntityIsWithinLifetimeOfExistingEntityReturnsTrueIfExistingEntityIsNotTerminatedAndNewEntitySourceDateIsAfterExistingEntitySourceDate() {

	lastTitleAttribute := models.Attribute{}.SetValue(test_values.TestValueObj)

	testEntity := models.Entity{}.SetTitle(test_values.TestValueObj2).SetSourceDate(test_values.TestValueObj2.GetDate())

	TestValue := managers.EntityManager{}.NewEntityIsWithinLifeTimeOfExistingEntity(testEntity, lastTitleAttribute, false)

	t.AssertEqual(testEntity.IsTerminated(), false)
	t.AssertEqual(testEntity.GetSourceDate().After(lastTitleAttribute.GetValue().Date), true)
	t.AssertEqual(TestValue, true)
}

/*
New entity is within lifetime of existing entity returns
true if existing entity is not terminated and new entity source date equals existing entity source date
 */
func (t *TestManagers) TestThatNewEntityIsWithinLifetimeOfExistingEntityReturnsTrueIfExistingEntityIsNotTerminatedAndNewEntitySourceDateEqualsExistingEntitySourceDate() {

	lastTitleAttribute := models.Attribute{}.SetValue(test_values.TestValueObj).SetValue(test_values.TestValueObj2).SetValue(test_values.TestValueObj3)

	testEntity := models.Entity{}.SetTitle(test_values.TestValueObj).SetSourceDate(test_values.TestValueObj.GetDate())

	TestValue := managers.EntityManager{}.NewEntityIsWithinLifeTimeOfExistingEntity(testEntity, lastTitleAttribute, false)

	t.AssertEqual(testEntity.IsTerminated(), false)
	t.AssertEqual(testEntity.GetSourceDate().Equal(lastTitleAttribute.GetValueByDate(test_values.TestValueObj.GetDate()).Date), true)
	t.AssertEqual(TestValue, true)
}

/*
New entity is within lifetime of existing entity returns
true if existing entity is terminated but new entity source date is between existing entity lifetime
 */
func (t *TestManagers) TestThatNewEntityIsWithinLifetimeOfExistingEntityReturnsTrueIfExistingEntityIsTerminatedAndNewEntitySourceDateIsWithinEntityLifetime() {

	lastTitleAttribute := models.Attribute{}.SetValue(test_values.TestValueObj).SetValue(test_values.TestValueObj2).SetValue(test_values.TestValueObj3)

	testEntity := models.Entity{}.SetTitle(test_values.TestValueObj).SetSourceDate(test_values.TestValueObj2.GetDate())

	TestValue := managers.EntityManager{}.NewEntityIsWithinLifeTimeOfExistingEntity(testEntity, lastTitleAttribute, true)

	t.AssertEqual(testEntity.GetSourceDate().After(lastTitleAttribute.GetValues()[0].GetDate()), true)
	t.AssertEqual(testEntity.GetSourceDate().Before(lastTitleAttribute.GetValue().GetDate()), true)
	t.AssertEqual(TestValue, true)
}

/*
New entity is within lifetime of existing entity returns
true if existing entity is terminated but new entity source date equals existing entity source date
 */
func (t *TestManagers) TestThatNewEntityIsWithinLifetimeOfExistingEntityReturnsTrueIfExistingEntityIsTerminatedButNewEntitySourceDateEqualsExistingEntitySourceDate() {

	lastTitleAttribute := models.Attribute{}.SetValue(test_values.TestValueObj).SetValue(test_values.TestValueObj2).SetValue(test_values.TestValueObj3)

	testEntity := models.Entity{}.SetTitle(test_values.TestValueObj).SetSourceDate(test_values.TestValueObj.GetDate())

	TestValue := managers.EntityManager{}.NewEntityIsWithinLifeTimeOfExistingEntity(testEntity, lastTitleAttribute, true)

	t.AssertEqual(testEntity.GetSourceDate().Equal(lastTitleAttribute.GetValues()[0].GetDate()), true)
	t.AssertEqual(TestValue, true)
}

/*
New entity is within lifetime of existing entity returns
false if existing entity is terminated and new entity source date is after existing entity termination date
 */
func (t *TestManagers) TestThatNewEntityIsWithinLifetimeOfExistingEntityReturnsFalseIfExistingEntityIsTerminatedAndNewEntitySourceDateAfterExistingEntityTerminationDate() {

	lastTitleAttribute := models.Attribute{}.SetValue(test_values.TestValueObj).SetValue(test_values.TestValueObj2)

	testEntity := models.Entity{}.SetTitle(test_values.TestValueObj3).SetSourceDate(test_values.TestValueObj3.GetDate())

	TestValue := managers.EntityManager{}.NewEntityIsWithinLifeTimeOfExistingEntity(testEntity, lastTitleAttribute, true)

	t.AssertEqual(testEntity.GetSourceDate().After(lastTitleAttribute.GetValue().GetDate()), true)
	t.AssertEqual(TestValue, false)
}

/*
New entity is within lifetime of existing entity returns
false if new entity source date is before existing entity source date
 */
func (t *TestManagers) TestThatNewEntityIsWithinLifetimeOfExistingEntityReturnsFalseIfNewEntitySourceDateIsBeforeExistingEntitySourceDate() {

	lastTitleAttribute := models.Attribute{}.SetValue(test_values.TestValueObj3).SetValue(test_values.TestValueObj2)
	newAttribute := models.Attribute{}.SetValue(test_values.TestValueObj)

	testEntity := models.Entity{}.SetTitle(test_values.TestValueObj).SetSourceDate(test_values.TestValueObj.GetDate())

	TestValue := managers.EntityManager{}.NewEntityIsWithinLifeTimeOfExistingEntity(testEntity, lastTitleAttribute, true)

	t.AssertEqual(newAttribute.GetValue().GetDate().Before(lastTitleAttribute.GetValue().GetDate()), true)
	t.AssertEqual(testEntity.GetSourceDate().Before(lastTitleAttribute.GetValues()[0].GetDate()), true)
	t.AssertEqual(TestValue, false)
}