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
//       "$ref": "#/definitions/SearchResult"
//
// responses:
//   '200':
//     description: user created/ modified
//     schema:
//         "$ref": "#/definitions/SearchResult"
//   '403':
//     description: input validation error
//     schema:
////       "$ref": "#/definitions/ErrorResponse"
//   '500':
//     description: server error
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
func (c UserController) Create() revel.Result {
	var (
		err  error
	)
	log.Println("create user request")

	password, _ := bcrypt.GenerateFromPassword([]byte("abc123"), 12)

	user := models.User{
		Name:     "admin",
		Role:     "admin",
		Email:    "umayangag@datafoundation.lk",
		Password: password,
	}

	//err = c.Params.BindJSON(&user)
	//if err != nil {
	//	log.Println("binding error:", err)
	//	c.Response.Status = 403
	//	return c.RenderJSON(controllers.BuildErrResponse(err))
	//}
	user, c.Response.Status, err = repositories.UserRepository{}.AddUser(user)
	if err != nil {
		log.Println("user create error:", err)
		c.Response.Status = 500
		return c.RenderJSON(controllers.BuildErrResponse(err,500))
	}
	return c.RenderJSON(user)

}
