package api

import (
	"GIG/app/controllers"
	"GIG/app/utility/normalizers"
	"GIG/app/utility/normalizers/locations"
	"GIG/app/utility/normalizers/names"
	"github.com/revel/revel"
)

type TestController struct {
	*revel.Controller
}

func (c TestController) NormalizeLocation() revel.Result {
	searchText := c.Params.Values.Get("searchText")
	if searchText == "" {
		c.Response.Status = 500
		return c.RenderJSON("searchText required")
	}
	result, err := locations.NormalizeLocation(searchText)
	if err != nil {
		errResp := controllers.BuildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(result)
}

func (c TestController) NormalizeName() revel.Result {
	searchText := c.Params.Values.Get("searchText")
	if searchText == "" {
		c.Response.Status = 500
		return c.RenderJSON("searchText required")
	}
	result, err := names.NormalizeName(searchText)
	if err != nil {
		errResp := controllers.BuildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(result)
}

func (c TestController) Normalize() revel.Result {
	searchText := c.Params.Values.Get("searchText")
	if searchText == "" {
		c.Response.Status = 500
		return c.RenderJSON("searchText required")
	}
	result, err := normalizers.Normalize(searchText)
	if err != nil {
		errResp := controllers.BuildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(result)
}
