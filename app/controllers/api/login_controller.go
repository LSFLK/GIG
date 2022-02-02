package api

import (
	"GIG-SDK/models"
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
//         "$ref": "#/definitions/SearchResult"
//   '403':
//     description: input validation error
//     schema:
////       "$ref": "#/definitions/ErrorResponse"
//   '500':
//     description: server error
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
func (c LoginController) Login() revel.Result {
	var credentials models.Login
	err := c.Params.BindJSON(&credentials)
	if err != nil {
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrResponse(err,403))
	}
	user, err := repositories.UserRepository{}.GetUserBy("name", credentials.Username)
	if err != nil {
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrResponse(errors.New("Invalid Credentials"),403))
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(credentials.Password)); err != nil {
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrResponse(errors.New("Invalid Credentials"),403))
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Email,
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	})

	secretKey, _:= revel.Config.String("app.secret")

	token, err := claims.SignedString([]byte(secretKey))


	c.Response.Out.Header().Set("Access-Control-Allow-Origin", "*")
	c.Response.Status = 200
	return c.RenderJSON(controllers.BuildSuccessResponse(token,200))

}
