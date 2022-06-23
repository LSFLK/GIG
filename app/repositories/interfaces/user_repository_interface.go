package interfaces

import "github.com/lsflk/gig-sdk/models"

type UserRepositoryInterface interface {
	AddUser(e models.User) (models.User, error)
	UpdateUser(e models.User) error
	GetUser(id string) (models.User, error)
	GetUserBy(attribute string, value string) (models.User, error)
	DeleteUser(user models.User) error
}
