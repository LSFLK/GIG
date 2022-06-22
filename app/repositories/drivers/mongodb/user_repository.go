package mongodb

import (
	"GIG/app/constants/database"
	"GIG/app/databases/mongodb"
	"GIG/app/repositories/interfaces"

	"github.com/lsflk/gig-sdk/models"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct {
	interfaces.UserRepositoryInterface
}

func (e UserRepository) newUserCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession(database.UserCollection)
}

/*
AddUser insert a new User into database and returns
last inserted user on success.
*/
func (e UserRepository) AddUser(user models.User) (models.User, error) {
	c := e.newUserCollection()
	defer c.Close()
	return user, c.Collection.Insert(user)
}

/*
GetUser - Get a User from database and returns
a models. User on success
*/
func (e UserRepository) GetUser(id string) (models.User, error) {
	var (
		user models.User
		err  error
	)

	c := e.newUserCollection()
	defer c.Close()

	err = c.Collection.Find(bson.M{"_id": id}).One(&user)
	return user, err
}

/*
GetUserBy - Get a User from database and returns
a models.User on success
*/
func (e UserRepository) GetUserBy(attribute string, value string) (models.User, error) {
	var (
		user models.User
		err  error
	)

	c := e.newUserCollection()
	defer c.Close()
	err = c.Collection.Find(bson.M{attribute: value}).Sort("-updated_at").One(&user)
	return user, err
}

/*
UpdateUser - update a User into database and returns
last nil on success.
*/
func (e UserRepository) UpdateUser(user models.User) error {
	c := e.newUserCollection()
	defer c.Close()

	err := c.Collection.Update(bson.M{
		"_id": user.GetId(),
	}, bson.M{
		"$set": user,
	})
	return err
}

/*
DeleteUser Delete User from database and returns
last nil on success.
*/
func (e UserRepository) DeleteUser(user models.User) error {
	c := e.newUserCollection()
	defer c.Close()

	err := c.Collection.Remove(bson.M{"_id": user.GetId()})
	return err
}
