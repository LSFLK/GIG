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
	"strings"
	"time"
)

type EntityController struct {
	*revel.Controller
}

// swagger:operation GET /search Entity search
//
// Search for entities by keywords and category
//
// This API allows search by category and key word searching to retrieve list of entities
//
// ---
// produces:
// - application/json
//
// parameters:
//
// - name: query
//   in: query
//   description: search keywords
//   required: false
//   type: string
//
// - name: categories
//   in: query
//   description: list of categories
//   required: false
//   type: array
//   items:
//     type: string
//   collectionFormat: csv
//
// - name: attributes
//   in: query
//   description: list of attributes to filter/ return all attributes if not provided
//   required: false
//   type: array
//   items:
//     type: string
//   collectionFormat: csv
//
// - name: page
//   in: query
//   description: page number of results
//   required: false
//   type: integer
//   format: int32
//
// - name: limit
//   in: query
//   description: maximum number of results to return
//   required: false
//   type: integer
//   format: int32
//
// responses:
//   '200':
//     description: search result
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/SearchResult"
//   '400':
//     description: input parameter validation error
//     schema:
////       "$ref": "#/definitions/ErrorResponse"
//   '500':
//     description: server error
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
func (c EntityController) Search() revel.Result {
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
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrResponse(errors.New("result limit is required")))
	}

	categoriesArray := libraries.ParseCategoriesString(categories)
	attributesArray := libraries.ParseCategoriesString(attributes)

	if searchKey == "" && categories == "" {
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrResponse(errors.New("search value or category is required")))
	}

	var responseArray []models.SearchResult
	entities, err = repositories.EntityRepository{}.GetEntities(searchKey, categoriesArray, limit, (page-1)*limit)
	if err != nil {
		log.Println(err)
		c.Response.Status = 500
		return c.RenderJSON(controllers.BuildErrResponse(err))
	}

	for _, element := range entities {
		responseArray = append(responseArray, models.SearchResult{}.ResultFrom(element, attributesArray))
	}
	c.Response.Status = 200
	return c.RenderJSON(responseArray)
}

// swagger:operation GET /get/{title} Entity show
//
// Return Entity
//
// This API allows key word searching to retrieve list of entities
//
// ---
// produces:
// - application/json
//
// parameters:
//
// - name: title
//   in: path
//   description: title of the entity
//   required: true
//   type: string
//
// - name: date
//   in: query
//   description: date to search the title for eg. 2006-01-02
//   required: false
//   type: date
//
// - name: attributes
//   in: query
//   description: list of attributes to filter/ return all attributes if not provided
//   required: false
//   type: array
//   items:
//     type: string
//   collectionFormat: csv
//
// - name: imageOnly
//   in: query
//   description: return only the default image.
//   required: false
//   type: boolean
//
// responses:
//   '200':
//     description: search result
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/SearchResult"
//   '202':
//     description: return default image path
//     schema:
//       type: string
//   '400':
//     description: input parameter validation error
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
//   '500':
//     description: server error
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
func (c EntityController) Show(title string) revel.Result {
	var (
		entity models.Entity
		err    error
	)
	log.Println("title", title)
	c.Response.Out.Header().Set("Access-Control-Allow-Origin", "*")

	if title == "" {
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrResponse(errors.New("invalid entity id format")))
	}
	dateParam := strings.Split(c.Params.Values.Get("date"), "T")[0]
	entityDate, dateError := time.Parse("2006-01-02", dateParam)
	attributes := c.Params.Values.Get("attributes")
	defaultImageOnly := c.Params.Values.Get("imageOnly")
	attributesArray := libraries.ParseCategoriesString(attributes)

	if dateError != nil || entityDate.IsZero() {
		entity, err = repositories.EntityRepository{}.GetEntityBy("title", title)
	} else {
		entity, err = repositories.EntityRepository{}.GetEntityByPreviousTitle(title, entityDate)
	}

	if err != nil {
		var normalizedName string
		normalizedName, err = repositories.EntityRepository{}.NormalizeEntityTitle(title)
		if err == nil {
			if dateError != nil || entityDate.IsZero() {
				entity, err = repositories.EntityRepository{}.GetEntityBy("title", normalizedName)
			} else {
				entity, err = repositories.EntityRepository{}.GetEntityByPreviousTitle(normalizedName, entityDate)
			}
		}
	}

	if err != nil {
		log.Println(err)
		c.Response.Status = 500
		return c.RenderJSON(controllers.BuildErrResponse(err))
	}

	// return only the default image url
	c.Response.Status = 202
	if defaultImageOnly == "true" {
		return c.RenderJSON(entity.ImageURL)
	}
	c.Response.Status = 200
	return c.RenderJSON(models.SearchResult{}.ResultFrom(entity, attributesArray))
}

func (c EntityController) CreateBatch() revel.Result {
	var (
		entities      []models.Entity
		savedEntities []models.Entity
	)
	log.Println("create entity batch request")
	err := c.Params.BindJSON(&entities)
	if err != nil {
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrResponse(err))
	}

	for _, e := range entities {
		entity, _, err := repositories.EntityRepository{}.AddEntity(e)
		if err != nil {
			c.Response.Status = 500
			return c.RenderJSON(controllers.BuildErrResponse(err))
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
	log.Println("create entity request")
	err = c.Params.BindJSON(&entity)
	if err != nil {
		log.Println("binding error:", err)
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrResponse(err))
	}
	entity, c.Response.Status, err = repositories.EntityRepository{}.AddEntity(entity)
	if err != nil {
		log.Println("entity create error:", err)
		c.Response.Status = 500
		return c.RenderJSON(controllers.BuildErrResponse(err))
	}
	return c.RenderJSON(entity)

}
