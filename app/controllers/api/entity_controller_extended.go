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
)

/**
 return list of entities linked inside a given entity
 */
func (c EntityController) GetEntityLinks(title string) revel.Result {
	var (
		entity        models.Entity
		responseArray []models.SearchResult
		err           error
	)

	limit, limitErr := strconv.Atoi(c.Params.Values.Get("limit"))
	page, pageErr := strconv.Atoi(c.Params.Values.Get("page"))
	attributes := c.Params.Values.Get("attributes")
	attributesArray := commons.ParseCategoriesString(attributes)
	if pageErr != nil || page < 1 {
		page = 1
	}

	c.Response.Out.Header().Set("Access-Control-Allow-Origin", "*")

	if limitErr != nil {
		errResp := controllers.BuildErrResponse(400, errors.New("result limit is required"), )
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	if title == "" {
		errResp := controllers.BuildErrResponse(400, errors.New("invalid entity id format"))
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	entity, err = repositories.EntityRepository{}.GetEntityBy("title", title)
	if err != nil {
		errResp := controllers.BuildErrResponse(500, err)
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	offset := (page - 1) * limit
	upperLimit := offset + limit
	if len(entity.GetLinks()) > offset {
		for i, linkTitle := range entity.GetLinks() {
			if i >= offset && i < upperLimit {
				linkedEntity, err := repositories.EntityRepository{}.GetEntityBy("title", linkTitle)
				if err != nil {
					fmt.Println(linkTitle, err)
				} else {
					responseArray = append(responseArray, models.SearchResult{}.ResultFrom(linkedEntity, attributesArray))
				}
			}
		}
	}

	c.Response.Status = 200
	return c.RenderJSON(responseArray)
}

/**
 return list of entities where a given entity is internally linked to it
 */
func (c EntityController) GetEntityRelations(title string) revel.Result {
	var (
		entities []models.Entity
		err      error
	)

	limit, limitErr := strconv.Atoi(c.Params.Values.Get("limit"))
	page, pageErr := strconv.Atoi(c.Params.Values.Get("page"))
	attributes := c.Params.Values.Get("attributes")
	attributesArray := commons.ParseCategoriesString(attributes)
	if pageErr != nil || page < 1 {
		page = 1
	}

	c.Response.Out.Header().Set("Access-Control-Allow-Origin", "*")

	if limitErr != nil {
		errResp := controllers.BuildErrResponse(400, errors.New("result limit is required"), )
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	if title == "" {
		errResp := controllers.BuildErrResponse(400, errors.New("invalid entity id format"))
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	entity, err := repositories.EntityRepository{}.GetEntityBy("title", title)
	if err != nil {
		errResp := controllers.BuildErrResponse(500, err)
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	entities, err = repositories.EntityRepository{}.GetRelatedEntities(entity, limit, (page-1)*limit)
	if err != nil {
		errResp := controllers.BuildErrResponse(500, err)
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	var responseArray []models.SearchResult
	for _, element := range entities {
		if element.GetTitle() != entity.GetTitle() { // exclude same entity from the result
			responseArray = append(responseArray, models.SearchResult{}.ResultFrom(element, attributesArray))
		}
	}
	c.Response.Status = 200
	return c.RenderJSON(responseArray)
}

/**
 terminate a list of entities or single entity
 */
func (c EntityController) TerminateEntities() revel.Result {
	var (
		entity   models.Entity
		entities []models.Entity
		err      error
	)
	fmt.Println("terminate entity request")
	err = c.Params.BindJSON(&entity)
	if err != nil {
		fmt.Println("binding error:", err)
		errResp := controllers.BuildErrResponse(403, err)
		c.Response.Status = 403
		return c.RenderJSON(errResp)
	}

	if entity.GetTitle() == "" && len(entity.GetCategories()) == 0 {
		errResp := controllers.BuildErrResponse(400, errors.New("title or category is required"), )
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	if entity.GetSourceDate().IsZero() || entity.GetSource() == "" {
		errResp := controllers.BuildErrResponse(400, errors.New("termination date and source is required"), )
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	if entity.GetTitle() != "" {
		existingEntity, err := repositories.EntityRepository{}.GetEntityBy("title", entity.GetTitle())
		if err != nil {
			fmt.Println(err)
			errResp := controllers.BuildErrResponse(500, err)
			c.Response.Status = 500
			return c.RenderJSON(errResp)
		}
		return c.RenderJSON(repositories.EntityRepository{}.TerminateEntity(existingEntity, entity.GetSource(), entity.GetSourceDate()))
	}

	entities, err = repositories.EntityRepository{}.GetEntities(entity.GetTitle(), entity.GetCategories(), 0, 0)
	if err != nil {
		fmt.Println(err)
		errResp := controllers.BuildErrResponse(500, err)
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	for _, result := range entities {
		repositories.EntityRepository{}.TerminateEntity(result, entity.GetSource(), entity.GetSourceDate())
	}
	c.Response.Status = 200
	return c.RenderJSON("entities terminated")
}
