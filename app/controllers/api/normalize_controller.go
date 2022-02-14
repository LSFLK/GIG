package api

import (
	"GIG/app/controllers"
	"GIG/app/repositories"
	"GIG/app/utilities/normalizers"
	"errors"
	"github.com/revel/revel"
)

type NormalizeController struct {
	*revel.Controller
}

// swagger:operation GET /normalize/location  Normalizer normalize-location
//
// Normalize a given location name to return a standard location name
//
// This API allows to normalize a given location name
//
// ---
// produces:
// - application/json
//
// parameters:
//
// - name: searchText
//   in: query
//   description: text to be normalized
//   required: true
//   type: string
//
// responses:
//   '200':
//     description: normalized text
//     schema:
//       type: object
//   '400':
//     description: input parameter validation error
//     schema:
////       "$ref": "#/definitions/Response"
//   '500':
//     description: server error
//     schema:
//       "$ref": "#/definitions/Response"
func (c NormalizeController) NormalizeLocation() revel.Result {
	searchText := c.Params.Values.Get("searchText")
	if searchText == "" {
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrorResponse(errors.New("searchText is required"),400))
	}
	result, err := normalizers.NormalizeLocation(searchText)
	if err != nil {
		c.Response.Status = 500
		return c.RenderJSON(controllers.BuildErrorResponse(err,500))
	}
	c.Response.Status = 200
	return c.RenderJSON(result)
}

// swagger:operation GET /normalize/name  Normalizer normalize-name
//
// Normalize a given entity title
//
// This API allows to normalize a given entity title
//
// ---
// produces:
// - application/json
//
// parameters:
//
// - name: searchText
//   in: query
//   description: text to be normalized
//   required: true
//   type: string
//
// responses:
//   '200':
//     description: normalized text
//     schema:
//       type: object
//   '400':
//     description: input parameter validation error
//     schema:
////       "$ref": "#/definitions/Response"
//   '500':
//     description: server error
//     schema:
//       "$ref": "#/definitions/Response"
func (c NormalizeController) NormalizeName() revel.Result {
	searchText := c.Params.Values.Get("searchText")
	if searchText == "" {
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrorResponse(errors.New("searchText is required"),400))
	}
	result, err := normalizers.NormalizeName(searchText)
	if err != nil {
		c.Response.Status = 500
		return c.RenderJSON(controllers.BuildErrorResponse(err,500))
	}
	c.Response.Status = 200
	return c.RenderJSON(result)
}

// swagger:operation GET /normalize  Normalizer normalize
//
// Normalize a given entity title
//
// This API allows to normalize a given entity title
//
// ---
// produces:
// - application/json
//
// parameters:
//
// - name: searchText
//   in: query
//   description: text to be normalized
//   required: true
//   type: string
//
// responses:
//   '200':
//     description: normalized text
//     schema:
//       type: object
//   '400':
//     description: input parameter validation error
//     schema:
////       "$ref": "#/definitions/Response"
//   '500':
//     description: server error
//     schema:
//       "$ref": "#/definitions/Response"
func (c NormalizeController) Normalize() revel.Result {
	searchText := c.Params.Values.Get("searchText")
	if searchText == "" {
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrorResponse(errors.New("searchText is required"), 400))
	}
	// try to get the normalized string from the system.
	normalizedName, err := repositories.EntityRepository{}.NormalizeEntityTitle(searchText)
	if err == nil {
		c.Response.Status = 200
		return c.RenderJSON(normalizedName)
	}

	c.Response.Status = 500
	return c.RenderJSON(controllers.BuildErrorResponse(err, 500))
}
