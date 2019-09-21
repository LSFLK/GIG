package api

import (
	"GIG/app/controllers"
	"GIG/app/models"
	"GIG/app/repositories/mongodb"
	"errors"
	"fmt"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type EntityController struct {
	*revel.Controller
}

func (c EntityController) Index() revel.Result {
	var (
		entities []models.Entity
		err      error
	)
	searchKey := c.Params.Values.Get("query")
	categories := c.Params.Values.Get("categories")
	var categoriesArray []string

	if strings.TrimSpace(categories) != "" {
		categoriesArray = strings.Split(categories, ",")
	}

	c.Response.Out.Header().Set("Access-Control-Allow-Origin", "*")

	if searchKey == "" && categories == "" {
		errResp := controllers.BuildErrResponse(400,errors.New("search value or category is required"), )
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	var responseArray []models.SearchResult
	entities, err = mongodb.GetEntities(searchKey, categoriesArray)
	if err != nil {
		fmt.Println(err)
		errResp := controllers.BuildErrResponse(500,err)
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	for _, element := range entities {
		responseArray = append(responseArray, models.SearchResult{}.ResultFrom(element))
	}
	c.Response.Status = 200
	return c.RenderJSON(responseArray)
}

func (c EntityController) Show(title string) revel.Result {
	var (
		entity models.Entity
		err    error
	)
	fmt.Println("title",title)
	c.Response.Out.Header().Set("Access-Control-Allow-Origin", "*")

	if title == "" {
		errResp := controllers.BuildErrResponse(400,errors.New("invalid entity id format"))
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	entity, err = mongodb.GetEntityBy("title", title)
	if err != nil {
		fmt.Println(err)
		errResp := controllers.BuildErrResponse(500,err )
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	c.Response.Status = 200
	return c.RenderJSON(entity)
}

func (c EntityController) CreateBatch() revel.Result {
	var (
		entities      []models.Entity
		savedEntities []models.Entity
	)
	fmt.Println("create entity batch request")
	err := c.Params.BindJSON(&entities)
	if err != nil {
		errResp := controllers.BuildErrResponse(403,err)
		c.Response.Status = 403
		return c.RenderJSON(errResp)
	}

	for _, e := range entities {
		entity, err := mongodb.AddEntity(e)
		if err != nil {
			errResp := controllers.BuildErrResponse(500,err)
			c.Response.Status = 500
			return c.RenderJSON(errResp)
		}
		savedEntities = append(savedEntities, entity)
	}

	c.Response.Status = 201
	return c.RenderJSON(savedEntities)
}

func (c EntityController) Create() revel.Result {
	var (
		entity models.Entity
		err    error
	)
	fmt.Println("create entity request")
	err = c.Params.BindJSON(&entity)
	if err != nil {
		fmt.Println("binding error:", err)
		errResp := controllers.BuildErrResponse(403,err)
		c.Response.Status = 403
		return c.RenderJSON(errResp)
	}
	entity, err = mongodb.AddEntity(entity)
	if err != nil {
		fmt.Println("entity create error:", err)
		errResp := controllers.BuildErrResponse(500,err)
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	fmt.Println("entity created", entity.Title)
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
		errResp := controllers.BuildErrResponse(400,err)
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	err = mongodb.UpdateEntity(entity)
	if err != nil {
		errResp := controllers.BuildErrResponse(500,err)
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
		errResp := controllers.BuildErrResponse(400,errors.New("invalid entity id format"))
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	entityID, err = controllers.ConvertToObjectIdHex(id)
	if err != nil {
		errResp := controllers.BuildErrResponse(400,errors.New("invalid entity id format"))
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	entity, err = mongodb.GetEntity(entityID)
	if err != nil {
		errResp := controllers.BuildErrResponse(500,err)
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	err = mongodb.DeleteEntity(entity)
	if err != nil {
		errResp := controllers.BuildErrResponse(500,err)
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 204
	return c.RenderJSON(nil)
}