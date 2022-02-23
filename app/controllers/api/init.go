package api

import (
	"GIG/app/repositories"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/revel/revel"
	"log"
	"net/http"
	"strings"
)

func AddLog(c *revel.Controller) revel.Result {
	log.Println("InterceptFunc Test.")
	return nil
}

var (
	errAuthHeaderNotFound = errors.New("authorization header not found")
	errInvalidTokenFormat = errors.New("token format is invalid")
)

func decodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		secretKey, _ := revel.Config.String("app.secret")
		return []byte(secretKey), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		fmt.Println("email and nbf:", claims["email"], claims["nbf"])
	} else {
		log.Println(err)
		return nil, err
	}
	return claims, nil
	// return claims["email"].(string), claims["nbf"].(string)
}

// Authenticate is and method will be called before any authenticate needed action.
// In order to valid the user.
func Authenticate(c *revel.Controller) revel.Result {
	tokenString, err := getTokenString(c, "Authorization")
	apiKey, keyErr := getTokenString(c, "ApiKey")

	if keyErr == nil {
		_, userErr := repositories.UserRepository{}.GetUserBy("apikey", apiKey)
		if userErr == nil && c.Name != "UserController" {
			return nil
		}
	}

	if err != nil {
		log.Println("get token/api key string failed")
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON("get token/api key string failed")
	}

	var claims jwt.MapClaims
	claims, err = decodeToken(tokenString)
	if err != nil {
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON("token decode error")
	}
	email, found := claims["iss"]
	if !found {
		log.Println(err)
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON("email not found in db")
	}
	user, err := repositories.UserRepository{}.GetUserBy("email", fmt.Sprintf("%v", email))
	if err != nil {
		log.Println(err)
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON("auth failed")
	}
	if user.Role != "admin" && c.Name == "UserController" {
		log.Println(err)
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON("auth failed")
	}
	
	log.Println("auth token success")
	return nil
}

func getTokenString(c *revel.Controller, headerName string) (tokenString string, err error) {
	authHeader := c.Request.Header.Get(headerName)
	if authHeader == "" {
		return "", errAuthHeaderNotFound
	}

	tokenSlice := strings.Split(authHeader, " ")
	if len(tokenSlice) != 2 {
		return "", errInvalidTokenFormat
	}
	tokenString = tokenSlice[1]
	return tokenString, nil

}

func init() {
	revel.InterceptFunc(Authenticate, revel.BEFORE, &UserController{})
	revel.InterceptFunc(Authenticate, revel.BEFORE, &PublisherController{})
	revel.InterceptFunc(Authenticate, revel.BEFORE, &EntityEditController{})
	revel.InterceptFunc(Authenticate, revel.BEFORE, &TokenValidationController{})
	revel.InterceptFunc(Authenticate, revel.BEFORE, &FileUploadController{})
}
