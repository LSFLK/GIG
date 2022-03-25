package api

import (
	"GIG/app/constants/error_messages"
	"GIG/app/constants/headers"
	"GIG/app/constants/info_messages"
	"GIG/app/constants/user_roles"
	"GIG/app/controllers"
	"GIG/app/repositories"
	"GIG/app/services/authentication"
	"github.com/lsflk/gig-sdk/models"
	"github.com/revel/revel"
	"github.com/tomogoma/go-typed-errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"log"
)

type ReaderController struct {
	*revel.Controller
}

// swagger:operation POST /register User register
//
// Create Reader
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
//       "$ref": "#/definitions/NewReader"
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
func (c ReaderController) Create() revel.Result {
	var (
		err       error
		newReader models.NewReader
	)
	log.Println(info_messages.UserCreateRequest)
	err = c.Params.BindJSON(&newReader)

	if err != nil {
		log.Println(error_messages.BindingError, err)
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrorResponse(err, 403))
	}

	if newReader.Email == "" || newReader.Password == "" {
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrorResponse(errors.New("Email and Password Required!"), 400))
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(newReader.Password), 12)
	apiKey, _ := bcrypt.GenerateFromPassword([]byte(newReader.Email), 12)

	user := models.User{
		Name:     newReader.Email,
		Role:     user_roles.Reader,
		Email:    newReader.Email,
		Password: password,
		ApiKey:   string(apiKey),
	}

	_, c.Response.Status, err = repositories.UserRepository{}.AddUser(user)
	if mgo.IsDup(err) {
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrorResponse(errors.New("Email Provided Already Exists!"), 500))
	}
	if err != nil {
		log.Println(error_messages.UserCreateError, err)
		c.Response.Status = 500
		return c.RenderJSON(controllers.BuildErrorResponse(errors.New("Error Registering User!"), 500))
	}
	userToken, err := authentication.CreateSignedUserToken(user)
	if err != nil {
		c.Response.Status = 500
		return c.RenderJSON(controllers.BuildErrorResponse(errors.New(error_messages.TokenSigningError), 403))
	}

	c.Response.Out.Header().Set(headers.AccessControlAllowOrigin, "*")
	c.Response.Status = 200
	return c.RenderJSON(controllers.BuildSuccessResponse(userToken, 200))

}
