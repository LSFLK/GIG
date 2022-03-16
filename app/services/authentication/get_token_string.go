package authentication

import (
	"GIG/app/constants/error_messages"
	"errors"
	"github.com/revel/revel"
	"strings"
)

func getTokenString(c *revel.Controller, headerName string) (tokenString string, err error) {
	authHeader := c.Request.Header.Get(headerName)
	if authHeader == "" {
		return "", errors.New(error_messages.AuthHeaderNotFound)
	}

	tokenSlice := strings.Split(authHeader, " ")
	if len(tokenSlice) != 2 {
		return "", errors.New(error_messages.InvalidTokenFormat)
	}
	tokenString = tokenSlice[1]
	return tokenString, nil

}
