package mongodb_official

import (
	"GIG/app/constants/database"
	"GIG/app/databases/mongodb_official"
	"GIG/app/repositories/interfaces"
	"github.com/lsflk/gig-sdk/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	interfaces.UserRepositoryInterface
}

func (e UserRepository) newUserCollection() *mongo.Collection {
	return mongodb_official.GetCollection(database.UserCollection)
}

/*
AddUser - insert a new User into database and returns
last inserted user on success.
*/
func (e UserRepository) AddUser(user models.User) (models.User, error) {
	c := e.newUserCollection()
	_, err := c.InsertOne(mongodb_official.Context, user)
	return user, err
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

	cursor := c.FindOne(mongodb_official.Context, bson.M{"_id": id})
	err = cursor.Decode(&user)
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
	findOptions := options.FindOne().SetSort(bson.D{{"updated_at", -1}})
	cursor := c.FindOne(mongodb_official.Context, bson.M{attribute: value}, findOptions)
	err = cursor.Decode(&user)
	return user, err
}

/*
UpdateUser - update a User into database and returns
last nil on success.
*/
func (e UserRepository) UpdateUser(user models.User) error {
	c := e.newUserCollection()

	filter := bson.D{{"_id", user.GetId()}}
	update := bson.D{{"$set", user}}
	_, err := c.UpdateOne(mongodb_official.Context, filter, update)
	return err
}

/*
DeleteUser - Delete User from database and returns
last nil on success.
*/
func (e UserRepository) DeleteUser(user models.User) error {
	c := e.newUserCollection()
	filter := bson.D{{"_id", user.GetId()}}
	_, err := c.DeleteOne(mongodb_official.Context, filter)
	return err
}
