package api

import (
	"GIG/app/controllers"
	"GIG/app/models"
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
		errResp := controllers.BuildErrResponse(400, errors.New("searchText is required"))
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}
	result, err := normalizers.NormalizeLocation(searchText)
	if err != nil {
		errResp := controllers.BuildErrResponse(500, err)
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(controllers.BuildResponse(200, result))
}

func (c NormalizeController) NormalizeName() revel.Result {
	searchText := c.Params.Values.Get("searchText")
	if searchText == "" {
		errResp := controllers.BuildErrResponse(400, errors.New("searchText is required"))
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}
	result, err := normalizers.NormalizeName(searchText)
	if err != nil {
		errResp := controllers.BuildErrResponse(500, err)
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(controllers.BuildResponse(200, result))
}

func (c NormalizeController) Normalize() revel.Result {
	searchText := c.Params.Values.Get("searchText")
	if searchText == "" {
		errResp := controllers.BuildErrResponse(400, errors.New("searchText is required"))
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}
	// try to get the normalized string from the system.
	normalizedName, err := repositories.NormalizedNameRepository{}.GetNormalizedNameBy("searchText", searchText)
	if err == nil {
		return c.RenderJSON(controllers.BuildResponse(200, normalizedName.NormalizedText))
	}
	//else call the service
	result, err := normalizers.Normalize(searchText)
	if err != nil || result == "" {
		errResp := controllers.BuildErrResponse(500, err)
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	//cache normalized string
	repositories.NormalizedNameRepository{}.AddNormalizedName(models.NormalizedName{NormalizedText: result})

	c.Response.Status = 200
	return c.RenderJSON(controllers.BuildResponse(200, result))
}
