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

func (c NormalizeController) NormalizeLocation() revel.Result {
	searchText := c.Params.Values.Get("searchText")
	if searchText == "" {
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrResponse( errors.New("searchText is required")))
	}
	result, err := normalizers.NormalizeLocation(searchText)
	if err != nil {
		c.Response.Status = 500
		return c.RenderJSON(controllers.BuildErrResponse( err))
	}
	c.Response.Status = 200
	return c.RenderJSON(result)
}

func (c NormalizeController) NormalizeName() revel.Result {
	searchText := c.Params.Values.Get("searchText")
	if searchText == "" {
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrResponse( errors.New("searchText is required")))
	}
	result, err := normalizers.NormalizeName(searchText)
	if err != nil {
		c.Response.Status = 500
		return c.RenderJSON( controllers.BuildErrResponse( err))
	}
	c.Response.Status = 200
	return c.RenderJSON(result)
}

func (c NormalizeController) Normalize() revel.Result {
	searchText := c.Params.Values.Get("searchText")
	if searchText == "" {
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrResponse( errors.New("searchText is required")))
	}
	// try to get the normalized string from the system.
	normalizedName, err := repositories.EntityRepository{}.NormalizeEntityTitle(searchText)
	if err == nil {
		c.Response.Status = 200
		return c.RenderJSON(normalizedName)
	}

	c.Response.Status = 500
	return c.RenderJSON(controllers.BuildErrResponse( err))
}
