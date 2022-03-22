package repositories

import (
	"GIG/app/repositories"
	"GIG/tests/app/test_values"
	"github.com/lsflk/gig-sdk/models"
)

/*
add entity works for new entity
 */
func (t *TestRepositories) TestThatAddEntityWorksForNewTitle() {

	testEntity := models.Entity{}.
		SetSourceSignature("trusted").
		SetTitle(test_values.TestValueObj.SetValueString("test entity for new title")).
		SetAttribute(test_values.TestAttributeKey, test_values.TestValueObj).AddCategory("TEST")

	savedEntity, err := repositories.EntityRepository{}.AddEntity(testEntity)

	t.AssertEqual(err, nil)
	t.AssertNotEqual(savedEntity.GetId().Hex(), "")

	deleteErr := repositories.EntityRepository{}.DeleteEntity(savedEntity)
	t.AssertEqual(deleteErr, nil)
}

/*
add entity works for existing entity with current title
 */
func (t *TestRepositories) TestThatAddEntityWorksForExistingEntityWithCurrentTitle() {

	testEntity := models.Entity{}.
		SetSourceSignature("trusted").
		SetTitle(test_values.TestValueObj.SetValueString("existing entity with current title")).
		SetAttribute(test_values.TestAttributeKey, test_values.TestValueObj).AddCategory("TEST")

	testEntity2 := models.Entity{}.
		SetSourceSignature("trusted").
		SetTitle(test_values.TestValueObj2.SetValueString("existing entity with current title")).
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
add entity works for existing entity with a previous title
 */
func (t *TestRepositories) TestThatAddEntityWorksForExistingEntityWithPreviousTitle() {

	testEntity := models.Entity{}.
		SetSourceSignature("trusted").
		SetTitle(test_values.TestValueObj.SetValueString("title value 1")).
		SetTitle(test_values.TestValueObj3.SetValueString("title value 3")).
		SetSourceDate(test_values.TestValueObj.GetDate()).
		SetAttribute(test_values.TestAttributeKey, test_values.TestValueObj).AddCategory("TEST")

	testEntity2 := models.Entity{}.
		SetSourceSignature("trusted").
		SetTitle(test_values.TestValueObj.SetValueString("title value 1")).
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

//TODO:
/*
get entity by previous title works for empty title
get entity by previous title works for same date source
 */

/*

get entity by previous title intermediate date
get entity by previous title future date
 */
