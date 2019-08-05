package api

import (
	"GIG/app/controllers"
	"GIG/app/utility/normalizers"
	"github.com/revel/revel"
)

type NormalizeController struct {
	*revel.Controller
}

func (c NormalizeController) NormalizeLocation() revel.Result {
	searchText := c.Params.Values.Get("searchText")
	if searchText == "" {
		c.Response.Status = 500
		return c.RenderJSON("searchText required")
	}
	result, err := normalizers.NormalizeLocation(searchText)
	if err != nil {
		errResp := controllers.BuildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(result)
}

func (c NormalizeController) NormalizeName() revel.Result {
	searchText := c.Params.Values.Get("searchText")
	if searchText == "" {
		c.Response.Status = 500
		return c.RenderJSON("searchText required")
	}
	result, err := normalizers.NormalizeName(searchText)
	if err != nil {
		errResp := controllers.BuildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(result)
}

func (c NormalizeController) Normalize() revel.Result {
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
