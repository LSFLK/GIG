package api

import (
	"GIG/app/controllers"
	"github.com/revel/revel"
)

type TokenValidationController struct {
	*revel.Controller
}

func (c TokenValidationController) ValidateToken() revel.Result {
	return c.RenderJSON(controllers.BuildSuccessResponse("token is valid", 200))
}
