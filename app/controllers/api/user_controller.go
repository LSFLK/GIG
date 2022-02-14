package api

import (
	"GIG-SDK/models"
	"GIG/app/controllers"
	"GIG/app/repositories"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserController struct {
	*revel.Controller
}

// swagger:operation POST /add User add
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
	log.Println("create user request")
	err = c.Params.BindJSON(&newUser)

	if err != nil {
		log.Println("binding error:", err)
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
		log.Println("user create error:", err)
		c.Response.Status = 500
		return c.RenderJSON(controllers.BuildErrorResponse(err, 500))
	}
	return c.RenderJSON(user)

}
