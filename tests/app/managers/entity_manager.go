package managers

import (
	"GIG/app/models"
	"GIG/app/utilities/managers"
)

/*
New entity title is within lifetime of existing entity returns false if entity is terminated
 */
func (t *TestManagers) TestThatNewEntityTitleIsWithinLifetimeOfExistingEntityReturnsFalseIfEntityIsTerminated() {
	testValue := managers.EntityManager{}.NewEntityTitleIsWithinLifetimeOfExistingEntity(models.Attribute{}, models.Attribute{}, true)

	t.AssertEqual(testValue, false)
}

/*
New entity title is within lifetime of existing entity returns true if new title date is after last title date
 */
func (t *TestManagers) TestThatNewEntityTitleIsWithinLifetimeOfExistingEntityReturnsTrueIfNewTitleDateIsAfterLastTitleDate() {

	lastTitleAttribute := models.Attribute{}.SetValue(testValueObj)
	newTitleAttribute := models.Attribute{}.SetValue(testValueObj2)

	testValue := managers.EntityManager{}.NewEntityTitleIsWithinLifetimeOfExistingEntity(newTitleAttribute, lastTitleAttribute, false)

	t.AssertEqual(newTitleAttribute.GetValue().GetDate().After(lastTitleAttribute.GetValue().GetDate()), true)
	t.AssertEqual(testValue, true)
}

/*
New entity title is within lifetime of existing entity returns false if new title date is before last title date
 */
func (t *TestManagers) TestThatNewEntityTitleIsWithinLifetimeOfExistingEntityReturnsFalseIfNewTitleDateIsBeforeLastTitleDate() {

	lastTitleAttribute := models.Attribute{}.SetValue(testValueObj2)
	newTitleAttribute := models.Attribute{}.SetValue(testValueObj)

	testValue := managers.EntityManager{}.NewEntityTitleIsWithinLifetimeOfExistingEntity(newTitleAttribute, lastTitleAttribute, false)

	t.AssertEqual(newTitleAttribute.GetValue().GetDate().Before(lastTitleAttribute.GetValue().GetDate()), true)
	t.AssertEqual(testValue, false)
}

/*
New entity title is within lifetime of existing entity returns false if new title date zero
 */
func (t *TestManagers) TestThatNewEntityTitleIsWithinLifetimeOfExistingEntityReturnsFalseIfNewTitleDateIsZero() {

	lastTitleAttribute := models.Attribute{}.SetValue(testValueObj2)
	newTitleAttribute := models.Attribute{}.SetValue(testValueObj0)

	testValue := managers.EntityManager{}.NewEntityTitleIsWithinLifetimeOfExistingEntity(newTitleAttribute, lastTitleAttribute, false)

	t.AssertEqual(newTitleAttribute.GetValue().GetDate().IsZero(), true)
	t.AssertEqual(testValue, false)
}

/*
New entity title is within lifetime of existing entity returns false if new title date equals last title date
 */
func (t *TestManagers) TestThatNewEntityTitleIsWithinLifetimeOfExistingEntityReturnsFalseIfNewTitleDateEqualsLastTitleDate() {

	lastTitleAttribute := models.Attribute{}.SetValue(testValueObj2)
	newTitleAttribute := models.Attribute{}.SetValue(testValueObj0)

	testValue := managers.EntityManager{}.NewEntityTitleIsWithinLifetimeOfExistingEntity(newTitleAttribute, lastTitleAttribute, false)

	t.AssertEqual(newTitleAttribute.GetValue().GetDate().IsZero(), true)
	t.AssertEqual(testValue, false)
}

/*
New entity is within lifetime of existing entity returns
true if existing entity is not terminated and new entity source date is after existing entity source date
 */
func (t *TestManagers) TestThatNewEntityIsWithinLifetimeOfExistingEntityReturnsTrueIfExistingEntityIsNotTerminatedAndNewEntitySourceDateIsAfterExistingEntitySourceDate() {

	lastTitleAttribute := models.Attribute{}.SetValue(testValueObj)

	testEntity := models.Entity{}.SetTitle(testValueObj2).SetSourceDate(testValueObj2.GetDate())

	testValue := managers.EntityManager{}.NewEntityIsWithinLifeTimeOfExistingEntity(testEntity, lastTitleAttribute, false)

	t.AssertEqual(testEntity.IsTerminated(), false)
	t.AssertEqual(testEntity.GetSourceDate().After(lastTitleAttribute.GetValue().Date), true)
	t.AssertEqual(testValue, true)
}

/*
New entity is within lifetime of existing entity returns
true if existing entity is not terminated and new entity source date equals existing entity source date
 */
func (t *TestManagers) TestThatNewEntityIsWithinLifetimeOfExistingEntityReturnsTrueIfExistingEntityIsNotTerminatedAndNewEntitySourceDateEqualsExistingEntitySourceDate() {

	lastTitleAttribute := models.Attribute{}.SetValue(testValueObj).SetValue(testValueObj2).SetValue(testValueObj3)

	testEntity := models.Entity{}.SetTitle(testValueObj).SetSourceDate(testValueObj.GetDate())

	testValue := managers.EntityManager{}.NewEntityIsWithinLifeTimeOfExistingEntity(testEntity, lastTitleAttribute, false)

	t.AssertEqual(testEntity.IsTerminated(), false)
	t.AssertEqual(testEntity.GetSourceDate().Equal(lastTitleAttribute.GetValueByDate(testValueObj.GetDate()).Date), true)
	t.AssertEqual(testValue, true)
}

/*
New entity is within lifetime of existing entity returns
true if existing entity is terminated but new entity source date is between existing entity lifetime
 */
func (t *TestManagers) TestThatNewEntityIsWithinLifetimeOfExistingEntityReturnsTrueIfExistingEntityIsTerminatedAndNewEntitySourceDateIsWithinEntityLifetime() {

	lastTitleAttribute := models.Attribute{}.SetValue(testValueObj).SetValue(testValueObj2).SetValue(testValueObj3)

	testEntity := models.Entity{}.SetTitle(testValueObj).SetSourceDate(testValueObj2.GetDate())

	testValue := managers.EntityManager{}.NewEntityIsWithinLifeTimeOfExistingEntity(testEntity, lastTitleAttribute, true)

	t.AssertEqual(testEntity.GetSourceDate().After(lastTitleAttribute.GetValues()[0].GetDate()), true)
	t.AssertEqual(testEntity.GetSourceDate().Before(lastTitleAttribute.GetValue().GetDate()), true)
	t.AssertEqual(testValue, true)
}

/*
New entity is within lifetime of existing entity returns
true if existing entity is terminated but new entity source date equals existing entity source date
 */
func (t *TestManagers) TestThatNewEntityIsWithinLifetimeOfExistingEntityReturnsTrueIfExistingEntityIsTerminatedButNewEntitySourceDateEqualsExistingEntitySourceDate() {

	lastTitleAttribute := models.Attribute{}.SetValue(testValueObj).SetValue(testValueObj2).SetValue(testValueObj3)

	testEntity := models.Entity{}.SetTitle(testValueObj).SetSourceDate(testValueObj.GetDate())

	testValue := managers.EntityManager{}.NewEntityIsWithinLifeTimeOfExistingEntity(testEntity, lastTitleAttribute, true)

	t.AssertEqual(testEntity.GetSourceDate().Equal(lastTitleAttribute.GetValues()[0].GetDate()), true)
	t.AssertEqual(testValue, true)
}

/*
New entity is within lifetime of existing entity returns
false if existing entity is terminated and new entity source date is after existing entity termination date
 */
func (t *TestManagers) TestThatNewEntityIsWithinLifetimeOfExistingEntityReturnsFalseIfExistingEntityIsTerminatedAndNewEntitySourceDateAfterExistingEntityTerminationDate() {

	lastTitleAttribute := models.Attribute{}.SetValue(testValueObj).SetValue(testValueObj2)

	testEntity := models.Entity{}.SetTitle(testValueObj3).SetSourceDate(testValueObj3.GetDate())

	testValue := managers.EntityManager{}.NewEntityIsWithinLifeTimeOfExistingEntity(testEntity, lastTitleAttribute, true)

	t.AssertEqual(testEntity.GetSourceDate().After(lastTitleAttribute.GetValue().GetDate()), true)
	t.AssertEqual(testValue, false)
}

/*
New entity is within lifetime of existing entity returns
false if new entity source date is before existing entity source date
 */
func (t *TestManagers) TestThatNewEntityIsWithinLifetimeOfExistingEntityReturnsFalseIfNewEntitySourceDateIsBeforeExistingEntitySourceDate() {

	lastTitleAttribute := models.Attribute{}.SetValue(testValueObj3).SetValue(testValueObj2)
	newAttribute := models.Attribute{}.SetValue(testValueObj)

	testEntity := models.Entity{}.SetTitle(testValueObj).SetSourceDate(testValueObj.GetDate())

	testValue := managers.EntityManager{}.NewEntityIsWithinLifeTimeOfExistingEntity(testEntity, lastTitleAttribute, true)

	t.AssertEqual(newAttribute.GetValue().GetDate().Before(lastTitleAttribute.GetValue().GetDate()), true)
	t.AssertEqual(testEntity.GetSourceDate().Before(lastTitleAttribute.GetValues()[0].GetDate()), true)
	t.AssertEqual(testValue, false)
}