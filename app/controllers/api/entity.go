package api

import (
	"GIG/app/models"
	"errors"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
)

type EntityController struct {
	*revel.Controller
}

func (c EntityController) Index() revel.Result {
	var (
		entities []models.Entity
		err      error
	)
	searchKey := c.Params.Values.Get("for")

	if searchKey == ""{
		return c.RenderJSON("search value is required")
	}

	entities, err = models.GetEntities(searchKey)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(entities)
}

func (c EntityController) Show(id string) revel.Result {
	var (
		entity   models.Entity
		err      error
		entityID bson.ObjectId
	)

	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid entity id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	entityID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid entity id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	entity, err = models.GetEntity(entityID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	c.Response.Status = 200
	return c.RenderJSON(entity)
}

func (c EntityController) Create() revel.Result {
	var (
		entity models.Entity
		err    error
	)
	err = c.Params.BindJSON(&entity)
	if err != nil {
		errResp := buildErrResponse(err, "403")
		c.Response.Status = 403
		return c.RenderJSON(errResp)
	}

	entity, err = models.AddEntity(entity)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 201
	return c.RenderJSON(entity)
}

func (c EntityController) Update() revel.Result {
	var (
		entity models.Entity
		err    error
	)
	err = c.Params.BindJSON(&entity)
	if err != nil {
		errResp := buildErrResponse(err, "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	err = entity.UpdateEntity()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	return c.RenderJSON(entity)
}

func (c EntityController) Delete(id string) revel.Result {
	var (
		err      error
		entity   models.Entity
		entityID bson.ObjectId
	)
	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid entity id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	entityID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid entity id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	entity, err = models.GetEntity(entityID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	err = entity.DeleteEntity()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 204
	return c.RenderJSON(nil)
}
