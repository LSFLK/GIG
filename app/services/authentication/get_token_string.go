package authentication

import (
	"GIG/app/constants/error_messages"
	"errors"
	"github.com/revel/revel"
	"strings"
)

func getTokenString(header *revel.RevelHeader, headerName string) (tokenString string, err error) {
	authHeader := header.Get(headerName)
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
