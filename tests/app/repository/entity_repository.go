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
		SetTitle(testValueObj).
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
		SetTitle(testValueObj).
		SetAttribute(testAttributeKey, testValueObj).AddCategory("TEST")

	testEntity2 := models.Entity{}.
		SetTitle(testValueObj2).
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

/*
get entity by previous title same date
get entity by previous title intermediate date
get entity by previous title future date
 */

/*
  terminate entity
doesn't terminate an already terminated entity
doesn;t terminate and entity if the termination source is older than latest title source date
 */
