package api

import (
	"GIG-SDK/models"
	"GIG/app/constants/error_messages"
	"GIG/app/constants/headers"
	"GIG/app/controllers"
	"GIG/app/repositories"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type LoginController struct {
	*revel.Controller
}

// swagger:operation POST /login
//
// Authenticate User
//
// This API allows to authenticate an user
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
//     description: login success
//     schema:
//         "$ref": "#/definitions/Login"
//   '403':
//     description: input validation error
//     schema:
////       "$ref": "#/definitions/Response"
//   '500':
//     description: server error
//     schema:
//       "$ref": "#/definitions/Response"
func (c LoginController) Login() revel.Result {
	var credentials models.Login
	err := c.Params.BindJSON(&credentials)
	if err != nil {
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrorResponse(err, 403))
	}
	user, err := repositories.UserRepository{}.GetUserBy("name", credentials.Username)
	if err != nil {
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrorResponse(errors.New(error_messages.InvalidLoginCredentials), 403))
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(credentials.Password)); err != nil {
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrorResponse(errors.New(error_messages.InvalidLoginCredentials), 403))
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Email,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Subject:   user.Role,
	})

	secretKey, _ := revel.Config.String("app.secret")

	token, err := claims.SignedString([]byte(secretKey))

	userToken := models.UserToken{Name: user.Name, Email: user.Email, Role: user.Role, Token: token}

	c.Response.Out.Header().Set(headers.AccessControlAllowOrigin, "*")
	c.Response.Status = 200
	return c.RenderJSON(controllers.BuildSuccessResponse(userToken, 200))

}
