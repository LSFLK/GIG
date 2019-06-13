package api

import (
	"GIG/app/controllers"
	"GIG/app/models"
	"errors"
	"fmt"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
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
		errResp := controllers.BuildErrResponse(errors.New("search value or category is required"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	var responseArray []models.SearchResult
	entities, err = models.GetEntities(searchKey, categoriesArray)

	for _, element := range entities {
		responseArray = append(responseArray, models.SearchResult{}.ResultFrom(element))
	}
	if err != nil {
		errResp := controllers.BuildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(responseArray)
}

func (c EntityController) Show(title string) revel.Result {
	var (
		entity models.Entity
		err    error
	)

	c.Response.Out.Header().Set("Access-Control-Allow-Origin", "*")

	if title == "" {
		errResp := controllers.BuildErrResponse(errors.New("invalid entity id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	entity, err = models.GetEntityBy("title", title)
	if err != nil {
		errResp := controllers.BuildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	entity = entity.EagerLoad()

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
		errResp := controllers.BuildErrResponse(err, "403")
		c.Response.Status = 403
		return c.RenderJSON(errResp)
	}

	entity.Title = strings.NewReplacer(
		"%", "",
		"/", "-",
		"~", "2",
	).Replace(entity.Title)

	existingEntity, _ := models.GetEntityBy("title", entity.Title)
	if entity.IsEqualTo(existingEntity) {
		c.Response.Status = 202
		return c.RenderJSON(existingEntity)
	}

	entity.ID = bson.NewObjectId()
	entity.UpdatedAt = time.Now()
	entity.CreatedAt = time.Now()

	entity, err = models.AddEntity(entity)
	if err != nil {
		fmt.Println(err)
		errResp := controllers.BuildErrResponse(err, "500")
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
		errResp := controllers.BuildErrResponse(err, "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	err = entity.UpdateEntity()
	if err != nil {
		errResp := controllers.BuildErrResponse(err, "500")
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
		errResp := controllers.BuildErrResponse(errors.New("invalid entity id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	entityID, err = controllers.ConvertToObjectIdHex(id)
	if err != nil {
		errResp := controllers.BuildErrResponse(errors.New("invalid entity id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	entity, err = models.GetEntity(entityID)
	if err != nil {
		errResp := controllers.BuildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	err = entity.DeleteEntity()
	if err != nil {
		errResp := controllers.BuildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 204
	return c.RenderJSON(nil)
}
