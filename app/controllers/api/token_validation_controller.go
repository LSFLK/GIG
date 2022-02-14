package api

import (
	"GIG/app/controllers"
	"github.com/revel/revel"
)

type TokenValidationController struct {
	*revel.Controller
}

// swagger:operation GET /validate
//
// Authenticate Validation of User
//
// This API allows to validate a token
//
// ---
// produces:
// - application/json
//
// responses:
//   '200':
//     description: login success
//     schema:
//         "$ref": "#/definitions/Response"
//   '403':
//     description: input validation error
//     schema:
////       "$ref": "#/definitions/Response"
//   '500':
//     description: server error
//     schema:
//       "$ref": "#/definitions/Response"
func (c TokenValidationController) ValidateToken() revel.Result {
	return c.RenderJSON(controllers.BuildSuccessResponse("token is valid", 200))
}
