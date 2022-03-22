package mongodb

import (
	"GIG/app/databases/mongodb"
	"github.com/lsflk/gig-sdk/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct {
}

func (e UserRepository) newUserCollection() *mongodb.Collection {
	c := mongodb.NewCollectionSession("users")
	userIndex := mgo.Index{
		Key:    []string{"name"},
		Name:   "userIndex",
		Unique: true,
	}
	emailIndex := mgo.Index{
		Key:    []string{"email"},
		Name:   "emailIndex",
		Unique: true,
	}
	c.Session.EnsureIndex(userIndex)
	c.Session.EnsureIndex(emailIndex)
	return c
}

/*
AddUser insert a new User into database and returns
last inserted user on success.
 */
func (e UserRepository) AddUser(user models.User) (models.User, error) {
	c := e.newUserCollection()
	defer c.Close()
	return user, c.Session.Insert(user)
}

/**
GetUser Get a User from database and returns
a models. User on success
 */
func (e UserRepository) GetUser(id bson.ObjectId) (models.User, error) {
	var (
		user models.User
		err    error
	)

	c := e.newUserCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"_id": id}).One(&user)
	return user, err
}

/**
GetUser Get a User from database and returns
a models.User on success
 */
func (e UserRepository) GetUserBy(attribute string, value string) (models.User, error) {
	var (
		user models.User
		err    error
	)

	c := e.newUserCollection()
	defer c.Close()
	err = c.Session.Find(bson.M{attribute: value}).Sort("-updated_at").One(&user)
	return user, err
}

/**
UpdateUser update a User into database and returns
last nil on success.
 */
func (e UserRepository) UpdateUser(user models.User) error {
	c := e.newUserCollection()
	defer c.Close()

	err := c.Session.Update(bson.M{
		"_id": user.GetId(),
	}, bson.M{
		"$set": user,
	})
	return err
}

/**
DeleteUser Delete User from database and returns
last nil on success.
 */
func (e UserRepository) DeleteUser(user models.User) error {
	c := e.newUserCollection()
	defer c.Close()

	err := c.Session.Remove(bson.M{"_id": user.GetId()})
	return err
}
