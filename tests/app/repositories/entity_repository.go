package repository

import (
	"GIG/app/models"
	"GIG/app/repositories"
)

/*
add entity works for new entity
 */
func (t *TestRepositories) TestThatAddEntityWorksForNewTitle() {

	testEntity := models.Entity{}.
		SetSourceSignature("trusted").
		SetTitle(testValueObj.SetValueString("test entity for new title")).
		SetAttribute(testAttributeKey, testValueObj).AddCategory("TEST")

	savedEntity, status, err := repositories.EntityRepository{}.AddEntity(testEntity)

	t.AssertEqual(err, nil)
	t.AssertEqual(status, 201)
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
		SetTitle(testValueObj.SetValueString("existing entity with current title")).
		SetAttribute(testAttributeKey, testValueObj).AddCategory("TEST")

	testEntity2 := models.Entity{}.
		SetSourceSignature("trusted").
		SetTitle(testValueObj2.SetValueString("existing entity with current title")).
		SetAttribute(testAttributeKey, testValueObj3).AddCategory("TEST")

	savedEntity, status, err := repositories.EntityRepository{}.AddEntity(testEntity)
	savedEntity2, status2, err2 := repositories.EntityRepository{}.AddEntity(testEntity2)

	savedAttribute, attributeErr := savedEntity2.GetAttribute(testAttributeKey)

	t.AssertEqual(err, nil)
	t.AssertEqual(err2, nil)
	t.AssertEqual(attributeErr, nil)
	t.AssertEqual(status, 201)
	t.AssertEqual(status2, 202)
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
		SetTitle(testValueObj.SetValueString("title value 1")).
		SetTitle(testValueObj3.SetValueString("title value 3")).
		SetSourceDate(testValueObj.GetDate()).
		SetAttribute(testAttributeKey, testValueObj).AddCategory("TEST")

	testEntity2 := models.Entity{}.
		SetSourceSignature("trusted").
		SetTitle(testValueObj.SetValueString("title value 1")).
		SetSourceDate(testValueObj2.GetDate()).
		SetAttribute(testAttributeKey, testValueObj2).AddCategory("TEST2")

	savedEntity, status, err := repositories.EntityRepository{}.AddEntity(testEntity)
	savedEntity2, status2, err2 := repositories.EntityRepository{}.AddEntity(testEntity2)

	savedAttribute, attributeErr := savedEntity2.GetAttribute(testAttributeKey)

	t.AssertEqual(err, nil)
	t.AssertEqual(err2, nil)
	t.AssertEqual(attributeErr, nil)
	t.AssertEqual(status, 201)
	t.AssertEqual(status2, 202)
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
