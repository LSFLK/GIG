package api

import (
	"GIG/app/services/authentication"
	"github.com/revel/revel"
)

func init() {
	revel.InterceptFunc(authentication.Authenticate, revel.BEFORE, &UserController{})
	revel.InterceptFunc(authentication.Authenticate, revel.BEFORE, &PublisherController{})
	revel.InterceptFunc(authentication.Authenticate, revel.BEFORE, &EntityEditController{})
	revel.InterceptFunc(authentication.Authenticate, revel.BEFORE, &TokenValidationController{})
	revel.InterceptFunc(authentication.Authenticate, revel.BEFORE, &FileUploadController{})
}
