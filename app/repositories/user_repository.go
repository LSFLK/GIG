package repositories

import (
	"github.com/lsflk/gig-sdk/models"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type iUserRepository interface {
	AddUser(e models.User) (models.User, error)
	UpdateUser(e models.User) error
	GetUser(id bson.ObjectId) (models.User, error)
	GetUserBy(attribute string, value string) (models.User, error)
	DeleteUser(user models.User) error
}

type UserRepository struct {
	iUserRepository
}

/*
AddUser insert a new User into database and returns
the user
*/
func (e UserRepository) AddUser(user models.User) (models.User, int, error) {

	log.Println("creating new user", user)
	user, err := repositoryHandler.userRepository.AddUser(user)
	return user, 201, err

}

/*
GetUser Get a User from database and returns
a models. User on success
*/
func (e UserRepository) GetUser(id bson.ObjectId) (models.User, error) {
	return repositoryHandler.userRepository.GetUser(id)
}

/*
GetUserBy - GetUser Get a User from database and returns
a models.User on success
*/
func (e UserRepository) GetUserBy(attribute string, value string) (models.User, error) {
	return repositoryHandler.userRepository.GetUserBy(attribute, value)
}

func (e UserRepository) DeleteUser(user models.User) error {
	return repositoryHandler.userRepository.DeleteUser(user)
}
