package authentication

import (
	"GIG-SDK/libraries"
	"GIG-SDK/models"
	"GIG/app/constants/error_messages"
	"GIG/app/constants/info_messages"
	"GIG/app/constants/user_roles"
	"GIG/app/repositories"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/revel/revel"
	"log"
	"net/http"
)

// Authenticate is and method will be called before any authenticate needed action.
// In order to valid the user.
func Authenticate(c *revel.Controller) revel.Result {

	user, authMethod, err := GetAuthUser(c.Request.Header)

	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(err.Error())
	}

	// if ApiKey exist and not requesting access to AdminControllers
	if authMethod == ApiKey && !libraries.StringInSlice(AdminControllers, c.Name) {
		return nil
	}

	if err != nil { // if Bearer token doesn't exist
		log.Println(error_messages.TokenApiKeyFailed)
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(error_messages.TokenApiKeyFailed)
	}

	if user.Role != user_roles.Admin && libraries.StringInSlice(AdminControllers, c.Name) { // Only admin users are allowed to access UserController
		log.Println(err)
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON(error_messages.InvalidLoginCredentials)
	}

	log.Println(info_messages.LoginSuccess, user.Email)
	return nil
}

func GetAuthUser(header *revel.RevelHeader) (models.User, string, error) {
	tokenString, err := getTokenString(header, AuthHeaderName)
	apiKey, keyErr := getTokenString(header, ApiKeyHeaderName)

	if keyErr == nil { // if ApiKey exist
		user, userErr := repositories.UserRepository{}.GetUserBy(ApiKey, apiKey)
		if userErr == nil {
			return user, ApiKey, nil
		}
	}

	if err != nil { // if Bearer token doesn't exist
		return models.User{}, Bearer, errors.New(error_messages.TokenApiKeyFailed)
	}

	var claims jwt.MapClaims
	claims, err = decodeToken(tokenString)
	if err != nil {
		return models.User{}, Bearer, errors.New(error_messages.TokenDecodeError)
	}
	email, found := claims["iss"]
	if !found {
		log.Println(err)
		return models.User{}, Bearer, err
	}
	user, err := repositories.UserRepository{}.GetUserBy(Email, fmt.Sprintf("%v", email))

	if err != nil {
		return models.User{}, Bearer, errors.New(error_messages.InvalidLoginCredentials)
	}

	return user, Bearer, nil
}
