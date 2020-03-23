package api

import (
	"GIG/app/controllers"
	"GIG/app/models"
	"GIG/app/repositories"
	"GIG/commons"
	"errors"
	"fmt"
	"github.com/revel/revel"
	"strconv"
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
	attributes := c.Params.Values.Get("attributes")
	limit, limitErr := strconv.Atoi(c.Params.Values.Get("limit"))
	page, pageErr := strconv.Atoi(c.Params.Values.Get("page"))
	if pageErr != nil || page < 1 {
		page = 1
	}

	c.Response.Out.Header().Set("Access-Control-Allow-Origin", "*")

	if limitErr != nil {
		errResp := controllers.BuildErrResponse(400, errors.New("result limit is required"), )
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	categoriesArray := commons.ParseCategoriesString(categories)
	attributesArray := commons.ParseCategoriesString(attributes)

	if searchKey == "" && categories == "" {
		errResp := controllers.BuildErrResponse(400, errors.New("search value or category is required"), )
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	var responseArray []models.SearchResult
	entities, err = repositories.EntityRepository{}.GetEntities(searchKey, categoriesArray, limit, (page-1)*limit)
	if err != nil {
		fmt.Println(err)
		errResp := controllers.BuildErrResponse(500, err)
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	for _, element := range entities {
		responseArray = append(responseArray, models.SearchResult{}.ResultFrom(element, attributesArray))
	}
	c.Response.Status = 200
	return c.RenderJSON(responseArray)
}

func (c EntityController) Show(title string) revel.Result {
	var (
		entity models.Entity
		err    error
	)
	fmt.Println("title", title)
	c.Response.Out.Header().Set("Access-Control-Allow-Origin", "*")

	if title == "" {
		errResp := controllers.BuildErrResponse(400, errors.New("invalid entity id format"))
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}
	entityDate, dateError := time.Parse("2006-1-2", c.Params.Values.Get("date"))
	attributes := c.Params.Values.Get("attributes")
	attributesArray := commons.ParseCategoriesString(attributes)

	if dateError != nil || entityDate.IsZero() {
		entity, err = repositories.EntityRepository{}.GetEntityBy("title", title)
	} else {
		entity, err = repositories.EntityRepository{}.GetEntityByPreviousTitle(title, entityDate)
	}

	if err != nil {
		fmt.Println(err)
		errResp := controllers.BuildErrResponse(500, err)
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	c.Response.Status = 200
	return c.RenderJSON(models.SearchResult{}.ResultFrom(entity, attributesArray))
}

func (c EntityController) CreateBatch() revel.Result {
	var (
		entities      []models.Entity
		savedEntities []models.Entity
	)
	fmt.Println("create entity batch request")
	err := c.Params.BindJSON(&entities)
	if err != nil {
		errResp := controllers.BuildErrResponse(403, err)
		c.Response.Status = 403
		return c.RenderJSON(errResp)
	}

	for _, e := range entities {
		entity, _, err := repositories.EntityRepository{}.AddEntity(e)
		if err != nil {
			errResp := controllers.BuildErrResponse(500, err)
			c.Response.Status = 500
			return c.RenderJSON(errResp)
		}
		savedEntities = append(savedEntities, entity)
	}

	c.Response.Status = 200
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
		errResp := controllers.BuildErrResponse(403, err)
		c.Response.Status = 403
		return c.RenderJSON(errResp)
	}
	entity, c.Response.Status, err = repositories.EntityRepository{}.AddEntity(entity)
	if err != nil {
		fmt.Println("entity create error:", err)
		errResp := controllers.BuildErrResponse(500, err)
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	return c.RenderJSON(entity)

}
