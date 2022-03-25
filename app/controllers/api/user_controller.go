package api

import (
	"GIG/app/constants/error_messages"
	"GIG/app/constants/info_messages"
	"GIG/app/controllers"
	"GIG/app/repositories"
	"github.com/lsflk/gig-sdk/models"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserController struct {
	*revel.Controller
}

// swagger:operation POST /user/add User add
//
// Create User
//
// This API allows to create/ modify a new/ existing user
//
// ---
// produces:
// - application/json
//
// parameters:
//
// - name: user
//   in: body
//   description: user object
//   required: true
//   schema:
//       "$ref": "#/definitions/NewUser"
//
// security:
//   - Bearer: []
// responses:
//   '200':
//     description: user created/ modified
//     schema:
//         "$ref": "#/definitions/User"
//   '403':
//     description: input validation error
//     schema:
////       "$ref": "#/definitions/Response"
//   '500':
//     description: server error
//     schema:
//       "$ref": "#/definitions/Response"
func (c UserController) Create() revel.Result {
	var (
		err     error
		newUser models.NewUser
	)
	log.Println(info_messages.UserCreateRequest)
	err = c.Params.BindJSON(&newUser)

	if err != nil {
		log.Println(error_messages.BindingError, err)
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrorResponse(err, 403))
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 12)
	apiKey, _ := bcrypt.GenerateFromPassword([]byte(newUser.Email), 12)

	user := models.User{
		Name:     newUser.Name,
		Role:     newUser.Role,
		Email:    newUser.Email,
		Password: password,
		ApiKey:   string(apiKey),
	}

	_, c.Response.Status, err = repositories.UserRepository{}.AddUser(user)
	if err != nil {
		log.Println(error_messages.UserCreateError, err)
		c.Response.Status = 500
		return c.RenderJSON(controllers.BuildErrorResponse(err, 500))
	}
	return c.RenderJSON(user)

}
