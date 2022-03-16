package authentication

import (
	"GIG-SDK/libraries"
	"GIG/app/constants/error_messages"
	"GIG/app/constants/info_messages"
	"GIG/app/constants/user_roles"
	"GIG/app/repositories"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/revel/revel"
	"log"
	"net/http"
)

// Authenticate is and method will be called before any authenticate needed action.
// In order to valid the user.
func Authenticate(c *revel.Controller) revel.Result {
	tokenString, err := getTokenString(c, AuthHeaderName)
	apiKey, keyErr := getTokenString(c, ApiKeyHeaderName)

	if keyErr == nil { // if ApiKey exist
		_, userErr := repositories.UserRepository{}.GetUserBy("apikey", apiKey)
		if userErr == nil && libraries.StringInSlice(AdminControllers, c.Name) { // Do not allow access to UserController using ApiKeys
			return nil
		}
	}

	if err != nil { // if Bearer token doesn't exist
		log.Println(error_messages.TokenApiKeyFailed)
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(error_messages.TokenApiKeyFailed)
	}

	var claims jwt.MapClaims
	claims, err = decodeToken(tokenString)
	if err != nil {
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON(error_messages.TokenDecodeError)
	}
	email, found := claims["iss"]
	if !found {
		log.Println(err)
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(error_messages.InvalidLoginCredentials)
	}
	user, err := repositories.UserRepository{}.GetUserBy("email", fmt.Sprintf("%v", email))
	if err != nil {
		log.Println(err)
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON(error_messages.InvalidLoginCredentials)
	}
	if user.Role != user_roles.Admin && libraries.StringInSlice(AdminControllers, c.Name) { // Only admin users are allowed to access UserController
		log.Println(err)
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON(error_messages.InvalidLoginCredentials)
	}

	log.Println(info_messages.LoginSuccess, email)
	return nil
}
