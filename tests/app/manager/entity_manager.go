package manager

import (
	"GIG/app/models"
	"GIG/app/repositories"
)

/*
New entity title is within lifetime of existing entity returns false if entity is terminated
 */
func (t *TestManagers) TestThatNewEntityTitleIsWithinLifetimeOfExistingEntityReturnsFalseIfEntityIsTerminated() {
	testValue := repositories.NewEntityTitleIsWithinLifetimeOfExistingEntity(models.Attribute{}, models.Attribute{}, true)
	t.AssertEqual(testValue, false)
}

/*
New entity title is within lifetime of existing entity returns true if new title date is after last title date
 */
func (t *TestManagers) TestThatNewEntityTitleIsWithinLifetimeOfExistingEntityReturnsTrueIfNewTitleDateIsAfterLastTitleDate() {

	lastTitleAttribute := models.Attribute{}.SetValue(testValueObj)
	newTitleAttribute := models.Attribute{}.SetValue(testValueObj2)

	testValue := repositories.NewEntityTitleIsWithinLifetimeOfExistingEntity(newTitleAttribute, lastTitleAttribute, false)
	t.AssertEqual(newTitleAttribute.GetValue().GetDate().After(lastTitleAttribute.GetValue().GetDate()), true)
	t.AssertEqual(testValue, true)
}

/*
New entity title is within lifetime of existing entity returns false if new title date is before last title date
 */
func (t *TestManagers) TestThatNewEntityTitleIsWithinLifetimeOfExistingEntityReturnsFalseIfNewTitleDateIsBeforeLastTitleDate() {

	lastTitleAttribute := models.Attribute{}.SetValue(testValueObj2)
	newTitleAttribute := models.Attribute{}.SetValue(testValueObj)

	testValue := repositories.NewEntityTitleIsWithinLifetimeOfExistingEntity(newTitleAttribute, lastTitleAttribute, false)
	t.AssertEqual(newTitleAttribute.GetValue().GetDate().Before(lastTitleAttribute.GetValue().GetDate()), true)
	t.AssertEqual(testValue, false)
}

/*
New entity title is within lifetime of existing entity returns false if new title date zero
 */
func (t *TestManagers) TestThatNewEntityTitleIsWithinLifetimeOfExistingEntityReturnsFalseIfNewTitleDateIsZero() {

	lastTitleAttribute := models.Attribute{}.SetValue(testValueObj2)
	newTitleAttribute := models.Attribute{}.SetValue(testValueObj0)

	testValue := repositories.NewEntityTitleIsWithinLifetimeOfExistingEntity(newTitleAttribute, lastTitleAttribute, false)
	t.AssertEqual(newTitleAttribute.GetValue().GetDate().IsZero(), true)
	t.AssertEqual(testValue, false)
}

/*
New entity title is within lifetime of existing entity returns false if new title date equals last title date
 */
func (t *TestManagers) TestThatNewEntityTitleIsWithinLifetimeOfExistingEntityReturnsFalseIfNewTitleDateEqualsLastTitleDate() {

	lastTitleAttribute := models.Attribute{}.SetValue(testValueObj2)
	newTitleAttribute := models.Attribute{}.SetValue(testValueObj0)

	testValue := repositories.NewEntityTitleIsWithinLifetimeOfExistingEntity(newTitleAttribute, lastTitleAttribute, false)
	t.AssertEqual(newTitleAttribute.GetValue().GetDate().IsZero(), true)
	t.AssertEqual(testValue, false)
}