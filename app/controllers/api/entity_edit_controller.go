package api

import (
	"GIG-SDK/models"
	"GIG/app/controllers"
	"GIG/app/repositories"
	"errors"
	"github.com/revel/revel"
	"log"
)

type EntityEditController struct {
	*revel.Controller
}

// swagger:operation POST /add-batch Entity add-batch
//
// Create a Set of Entities
//
// This API allows to create/ modify a new/ set of entities
//
// ---
// produces:
// - application/json
//
// parameters:
//
// - name: entity
//   in: body
//   description: entity object array
//   required: true
//   schema:
//       type: array
//       items:
//         "$ref": "#/definitions/SearchResult"
//
// responses:
//   '200':
//     description: entity created/ modified
//     schema:
//         "$ref": "#/definitions/SearchResult"
//   '403':
//     description: input validation error
//     schema:
////       "$ref": "#/definitions/ErrorResponse"
//   '500':
//     description: server error
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
func (c EntityEditController) CreateBatch() revel.Result {
	var (
		entities      []models.Entity
		savedEntities []models.Entity
	)
	log.Println("create entity batch request")
	err := c.Params.BindJSON(&entities)
	if err != nil {
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrResponse(err, 403))
	}

	for _, e := range entities {
		entity, _, err := repositories.EntityRepository{}.AddEntity(e)
		if err != nil {
			c.Response.Status = 500
			return c.RenderJSON(controllers.BuildErrResponse(err, 500))
		}
		savedEntities = append(savedEntities, entity)
	}

	c.Response.Status = 200
	return c.RenderJSON(savedEntities)
}

// swagger:operation POST /add Entity add
//
// Create Entity
//
// This API allows to create/ modify a new/ existing entity
//
// ---
// produces:
// - application/json
//
// parameters:
//
// - name: entity
//   in: body
//   description: entity object
//   required: true
//   schema:
//       "$ref": "#/definitions/SearchResult"
//
// responses:
//   '200':
//     description: entity created/ modified
//     schema:
//         "$ref": "#/definitions/SearchResult"
//   '403':
//     description: input validation error
//     schema:
////       "$ref": "#/definitions/ErrorResponse"
//   '500':
//     description: server error
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
func (c EntityEditController) Create() revel.Result {
	var (
		entity models.Entity
		err    error
	)
	log.Println("create entity request")
	err = c.Params.BindJSON(&entity)
	if err != nil {
		log.Println("binding error:", err)
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrResponse(err, 403))
	}
	entity, c.Response.Status, err = repositories.EntityRepository{}.AddEntity(entity)
	if err != nil {
		log.Println("entity create error:", err)
		c.Response.Status = 500
		return c.RenderJSON(controllers.BuildErrResponse(err, 500))
	}
	return c.RenderJSON(entity)

}

// swagger:operation POST /terminate Entity terminate
//
// Terminate Entities
//
// This API allows to terminate the lifetime of an existing entity. Include entity title to terminate specific entity or include categories to terminate set of entities by category.
// source date and source attributes are required*.
//
// ---
// produces:
// - application/json
//
// parameters:
//
// - name: entity
//   in: body
//   description: entity prototype
//   required: true
//   schema:
//       "$ref": "#/definitions/SearchResult"
//
// responses:
//   '200':
//     description: entity created/ modified
//     schema:
//         "$ref": "#/definitions/SearchResult"
//   '403':
//     description: input validation error
//     schema:
////       "$ref": "#/definitions/ErrorResponse"
//   '500':
//     description: server error
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
func (c EntityEditController) TerminateEntities() revel.Result {
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
		return c.RenderJSON(controllers.BuildErrResponse(err, 403))
	}

	if entity.GetTitle() == "" && len(entity.GetCategories()) == 0 {
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrResponse(errors.New("title or category is required"), 400))
	}

	if entity.GetSourceDate().IsZero() || entity.GetSource() == "" {
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrResponse(errors.New("termination date and source is required"), 400))
	}

	if entity.GetTitle() != "" {
		existingEntity, err := repositories.EntityRepository{}.GetEntityBy("title", entity.GetTitle())
		if err != nil {
			log.Println(err)
			c.Response.Status = 500
			return c.RenderJSON(controllers.BuildErrResponse(err, 500))
		}
		return c.RenderJSON(repositories.EntityRepository{}.TerminateEntity(existingEntity, entity.GetSource(), entity.GetSourceDate()))
	}

	entities, err = repositories.EntityRepository{}.GetEntities(entity.GetTitle(), entity.GetCategories(), 0, 0)
	if err != nil {
		log.Println(err)
		c.Response.Status = 500
		return c.RenderJSON(controllers.BuildErrResponse(err, 500))
	}

	for _, result := range entities {
		repositories.EntityRepository{}.TerminateEntity(result, entity.GetSource(), entity.GetSourceDate())
	}
	c.Response.Status = 200
	return c.RenderJSON("entities terminated")
}
