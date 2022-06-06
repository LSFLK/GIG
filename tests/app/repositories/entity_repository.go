package repositories

import (
	"GIG/app/repositories"
	"GIG/tests/app/test_values"
	"github.com/lsflk/gig-sdk/models"
)

/*
TestThatAddEntityWorksForNewTitle
add entity works for new entity
*/
func (t *TestRepositories) TestThatAddEntityWorksForNewTitle() {

	testValueObj0NewTitle := test_values.TestValueObj
	testEntity := *new(models.Entity).
		SetSourceSignature("trusted").
		SetTitle(*testValueObj0NewTitle.SetValueString("test entity for new title")).
		SetAttribute(test_values.TestAttributeKey, test_values.TestValueObj).AddCategory("TEST")

	savedEntity, err := repositories.EntityRepository{}.AddEntity(testEntity)

	t.AssertEqual(err, nil)
	t.AssertNotEqual(savedEntity.GetId().Hex(), "")

	deleteErr := repositories.EntityRepository{}.DeleteEntity(savedEntity)
	t.AssertEqual(deleteErr, nil)
}

/*
TestThatAddEntityWorksForExistingEntityWithCurrentTitle
add entity works for existing entity with current title
*/
func (t *TestRepositories) TestThatAddEntityWorksForExistingEntityWithCurrentTitle() {
	TestValueObjWithCurrentTitle := test_values.TestValueObj
	TestValueObj2WithCurrentTitle := test_values.TestValueObj2
	testEntity := *new(models.Entity).
		SetSourceSignature("trusted").
		SetTitle(*TestValueObjWithCurrentTitle.SetValueString("existing entity with current title")).
		SetAttribute(test_values.TestAttributeKey, test_values.TestValueObj).AddCategory("TEST")

	testEntity2 := *new(models.Entity).
		SetSourceSignature("trusted").
		SetTitle(*TestValueObj2WithCurrentTitle.SetValueString("existing entity with current title")).
		SetAttribute(test_values.TestAttributeKey, test_values.TestValueObj3).AddCategory("TEST")

	savedEntity, err := repositories.EntityRepository{}.AddEntity(testEntity)
	savedEntity2, err2 := repositories.EntityRepository{}.AddEntity(testEntity2)

	savedAttribute, attributeErr := savedEntity2.GetAttribute(test_values.TestAttributeKey)

	t.AssertEqual(err, nil)
	t.AssertEqual(err2, nil)
	t.AssertEqual(attributeErr, nil)
	t.AssertEqual(savedEntity.GetId(), savedEntity2.GetId())
	t.AssertEqual(len(savedAttribute.GetValues()), 2)
	t.AssertNotEqual(savedEntity.GetId().Hex(), "")
	t.AssertNotEqual(savedEntity2.GetId().Hex(), "")

	deleteErr := repositories.EntityRepository{}.DeleteEntity(savedEntity2)
	t.AssertEqual(deleteErr, nil)
}

/*
TestThatAddEntityWorksForExistingEntityWithPreviousTitle
add entity works for existing entity with a previous title
*/
func (t *TestRepositories) TestThatAddEntityWorksForExistingEntityWithPreviousTitle() {
	TestValueObjValue1 := test_values.TestValueObj
	TestValueObj3Value3 := test_values.TestValueObj3
	testEntity := *new(models.Entity).
		SetSourceSignature("trusted").
		SetTitle(*TestValueObjValue1.SetValueString("title value 1")).
		SetTitle(*TestValueObj3Value3.SetValueString("title value 3")).
		SetSourceDate(test_values.TestValueObj.GetDate()).
		SetAttribute(test_values.TestAttributeKey, test_values.TestValueObj).AddCategory("TEST")

	testEntity2 := *new(models.Entity).
		SetSourceSignature("trusted").
		SetTitle(*TestValueObjValue1.SetValueString("title value 1")).
		SetSourceDate(test_values.TestValueObj2.GetDate()).
		SetAttribute(test_values.TestAttributeKey, test_values.TestValueObj2).AddCategory("TEST2")

	savedEntity, err := repositories.EntityRepository{}.AddEntity(testEntity)
	savedEntity2, err2 := repositories.EntityRepository{}.AddEntity(testEntity2)

	savedAttribute, attributeErr := savedEntity2.GetAttribute(test_values.TestAttributeKey)

	t.AssertEqual(err, nil)
	t.AssertEqual(err2, nil)
	t.AssertEqual(attributeErr, nil)
	t.AssertEqual(savedEntity.GetId(), savedEntity2.GetId())
	t.AssertEqual(len(savedAttribute.GetValues()), 2)
	t.AssertNotEqual(savedEntity.GetId().Hex(), "")
	t.AssertNotEqual(savedEntity2.GetId().Hex(), "")

	deleteErr := repositories.EntityRepository{}.DeleteEntity(savedEntity2)
	t.AssertEqual(deleteErr, nil)
}

/*
TestThatAddEntityWorksForExistingEntityWithSameTitleAndSourceDate
add entity doesn't create duplicate attribute for same source date and value
*/
func (t *TestRepositories) TestThatAddEntityWorksForExistingEntityWithSameTitleAndSourceDate() {

	testEntity := *new(models.Entity).
		SetSourceSignature("trusted").
		SetTitle(test_values.TestValueObj).
		SetSourceDate(test_values.TestValueObj.GetDate()).
		SetAttribute(test_values.TestAttributeKey, test_values.TestValueObj).AddCategory("TEST")

	testEntity2 := *new(models.Entity).
		SetSourceSignature("trusted").
		SetTitle(test_values.TestValueObj).
		SetSourceDate(test_values.TestValueObj.GetDate()).
		SetAttribute(test_values.TestAttributeKey, test_values.TestValueObj).AddCategory("TEST2")

	savedEntity, err := repositories.EntityRepository{}.AddEntity(testEntity)
	savedEntity2, err2 := repositories.EntityRepository{}.AddEntity(testEntity2)

	savedAttribute, attributeErr := savedEntity2.GetAttribute(test_values.TestAttributeKey)

	t.AssertEqual(err, nil)
	t.AssertEqual(err2, nil)
	t.AssertEqual(attributeErr, nil)
	t.AssertEqual(savedEntity.GetId(), savedEntity2.GetId())
	t.AssertEqual(len(savedAttribute.GetValues()), 1)
	t.AssertNotEqual(savedEntity.GetId().Hex(), "")
	t.AssertNotEqual(savedEntity2.GetId().Hex(), "")

	deleteErr := repositories.EntityRepository{}.DeleteEntity(savedEntity2)
	t.AssertEqual(deleteErr, nil)
}
