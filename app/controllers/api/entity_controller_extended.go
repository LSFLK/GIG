package api

import (
	"GIG-SDK/libraries"
	"GIG-SDK/models"
	"GIG/app/controllers"
	"GIG/app/repositories"
	"errors"
	"github.com/revel/revel"
	"log"
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
	attributesArray := libraries.ParseCategoriesString(attributes)
	if pageErr != nil || page < 1 {
		page = 1
	}

	c.Response.Out.Header().Set("Access-Control-Allow-Origin", "*")

	if limitErr != nil {
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrResponse(errors.New("result limit is required")))
	}

	if title == "" {
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrResponse(errors.New("invalid entity id format")))
	}

	entity, err = repositories.EntityRepository{}.GetEntityBy("title", title)
	if err != nil {
		c.Response.Status = 500
		return c.RenderJSON(controllers.BuildErrResponse(err))
	}

	offset := (page - 1) * limit
	upperLimit := offset + limit
	if len(entity.GetLinks()) > offset {
		for i, link := range entity.GetLinks() {
			if i >= offset && i < upperLimit {
				var (
					linkedEntity models.Entity
					err          error
				)
				if len(link.GetDates()) > 0 {
					linkedEntity, err = repositories.EntityRepository{}.GetEntityByPreviousTitle(link.GetTitle(), link.GetDates()[0])
				} else {
					linkedEntity, err = repositories.EntityRepository{}.GetEntityBy("title", link.GetTitle())
				}

				if err != nil {
					log.Println(link.GetTitle(), err)
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
	attributesArray := libraries.ParseCategoriesString(attributes)
	if pageErr != nil || page < 1 {
		page = 1
	}

	c.Response.Out.Header().Set("Access-Control-Allow-Origin", "*")

	if limitErr != nil {
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrResponse(errors.New("result limit is required")))
	}

	if title == "" {
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrResponse(errors.New("invalid entity id format")))
	}

	entity, err := repositories.EntityRepository{}.GetEntityBy("title", title)
	if err != nil {
		c.Response.Status = 500
		return c.RenderJSON(controllers.BuildErrResponse(err))
	}

	entities, err = repositories.EntityRepository{}.GetRelatedEntities(entity, limit, (page-1)*limit)
	if err != nil {
		c.Response.Status = 500
		return c.RenderJSON(controllers.BuildErrResponse(err))
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
	log.Println("terminate entity request")
	err = c.Params.BindJSON(&entity)
	if err != nil {
		log.Println("binding error:", err)
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrResponse(err))
	}

	if entity.GetTitle() == "" && len(entity.GetCategories()) == 0 {
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrResponse(errors.New("title or category is required")))
	}

	if entity.GetSourceDate().IsZero() || entity.GetSource() == "" {
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrResponse(errors.New("termination date and source is required")))
	}

	if entity.GetTitle() != "" {
		existingEntity, err := repositories.EntityRepository{}.GetEntityBy("title", entity.GetTitle())
		if err != nil {
			log.Println(err)
			c.Response.Status = 500
			return c.RenderJSON(controllers.BuildErrResponse(err))
		}
		return c.RenderJSON(repositories.EntityRepository{}.TerminateEntity(existingEntity, entity.GetSource(), entity.GetSourceDate()))
	}

	entities, err = repositories.EntityRepository{}.GetEntities(entity.GetTitle(), entity.GetCategories(), 0, 0)
	if err != nil {
		log.Println(err)
		c.Response.Status = 500
		return c.RenderJSON(controllers.BuildErrResponse(err))
	}

	for _, result := range entities {
		repositories.EntityRepository{}.TerminateEntity(result, entity.GetSource(), entity.GetSourceDate())
	}
	c.Response.Status = 200
	return c.RenderJSON("entities terminated")
}
