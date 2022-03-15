package api

import (
	"GIG-SDK/enums/ValueType"
	"GIG-SDK/models"
	"GIG/app/controllers"
	"GIG/app/repositories"
	"errors"
	"github.com/revel/revel"
	"log"
	"sync"
	"time"
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
// security:
//   - Bearer: []
//   - ApiKey: []
//
// responses:
//   '200':
//     description: entity created/ modified
//     schema:
//         "$ref": "#/definitions/SearchResult"
//   '403':
//     description: input validation error
//     schema:
////       "$ref": "#/definitions/Response"
//   '500':
//     description: server error
//     schema:
//       "$ref": "#/definitions/Response"
func (c EntityEditController) CreateBatch() revel.Result {
	var (
		entitiesList []models.Entity
	)
	log.Println("create entity batch request")
	err := c.Params.BindJSON(&entitiesList)
	if err != nil {
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrorResponse(err, 403))
	}

	go func(entities []models.Entity) {
		wg := &sync.WaitGroup{}

		for _, e := range entities {
			wg.Add(1)
			go func(entity models.Entity) {
				_, _, err := repositories.EntityRepository{}.AddEntity(entity)
				if err != nil {
					log.Println("entity creation error:", e)
				}
			}(e)

		}

		wg.Wait()
	}(entitiesList)

	c.Response.Status = 200
	return c.RenderJSON(controllers.BuildSuccessResponse("entity batch creation queued.", 200))
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
// security:
//   - Bearer: []
//   - ApiKey: []
//
// responses:
//   '200':
//     description: entity created/ modified
//     schema:
//         "$ref": "#/definitions/SearchResult"
//   '403':
//     description: input validation error
//     schema:
////       "$ref": "#/definitions/Response"
//   '500':
//     description: server error
//     schema:
//       "$ref": "#/definitions/Response"
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
		return c.RenderJSON(controllers.BuildErrorResponse(err, 403))
	}

	go func(newEntity models.Entity) {
		entity, c.Response.Status, err = repositories.EntityRepository{}.AddEntity(newEntity)
		if err != nil {
			log.Println("entity creation error:", err)
		}
	}(entity)

	return c.RenderJSON(controllers.BuildSuccessResponse("entity creation queued.", 200))

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
// security:
//   - Bearer: []
//   - ApiKey: []
//
// responses:
//   '200':
//     description: entity created/ modified
//     schema:
//         "$ref": "#/definitions/SearchResult"
//   '403':
//     description: input validation error
//     schema:
////       "$ref": "#/definitions/Response"
//   '500':
//     description: server error
//     schema:
//       "$ref": "#/definitions/Response"
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
		return c.RenderJSON(controllers.BuildErrorResponse(err, 403))
	}

	if entity.GetTitle() == "" && len(entity.GetCategories()) == 0 {
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrorResponse(errors.New("title or category is required"), 400))
	}

	if entity.GetSourceDate().IsZero() || entity.GetSource() == "" {
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrorResponse(errors.New("termination date and source is required"), 400))
	}

	go func(passedEntity models.Entity) {
		if passedEntity.GetTitle() != "" {
			existingEntity, err := repositories.EntityRepository{}.GetEntityBy("title", passedEntity.GetTitle())
			if err != nil {
				log.Println(err)
			}
			err = repositories.EntityRepository{}.TerminateEntity(existingEntity, passedEntity.GetSource(), passedEntity.GetSourceDate())
			if err != nil {
				log.Println(err)
			}
		}

		entities, err = repositories.EntityRepository{}.GetEntities(passedEntity.GetTitle(), passedEntity.GetCategories(), 0, 0)
		if err != nil {
			log.Println(err)
		}

		for _, result := range entities {
			err = repositories.EntityRepository{}.TerminateEntity(result, passedEntity.GetSource(), passedEntity.GetSourceDate())
			if err != nil {
				log.Println(err)
			}
		}
	}(entity)

	return c.RenderJSON(controllers.BuildSuccessResponse("entity termination queued.", 200))
}

// swagger:operation POST /delete Entity delete
//
// Delete Entity
//
// This API allows to delete existing entity
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
// security:
//   - Bearer: []
//   - ApiKey: []
//
// responses:
//   '200':
//     description: entity created/ modified
//     schema:
//         "$ref": "#/definitions/SearchResult"
//   '403':
//     description: input validation error
//     schema:
////       "$ref": "#/definitions/Response"
//   '500':
//     description: server error
//     schema:
//       "$ref": "#/definitions/Response"
func (c EntityEditController) DeleteEntity() revel.Result {
	var (
		entity models.Entity
		err    error
	)

	err = c.Params.BindJSON(&entity)
	if err != nil {
		log.Println("binding error:", err)
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrorResponse(err, 403))
	}
	log.Println("delete entity request", entity.Title)

	go func(entityToDelete models.Entity) {
		existingEntity, err := repositories.EntityRepository{}.GetEntityByPreviousTitle(entityToDelete.Title, time.Now())
		if err != nil {
			log.Println("error finding entity:", err)
		}

		err = repositories.EntityRepository{}.DeleteEntity(existingEntity)
		if err != nil {
			log.Println("error deleting entity:", err)
		}
	}(entity)

	return c.RenderJSON(controllers.BuildSuccessResponse("entity deletion queued.", 200))
}

// swagger:operation POST /update Entity update
//
// Update Entity
//
// This API allows to modify existing entity
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
// security:
//   - Bearer: []
//   - ApiKey: []
//
// responses:
//   '200':
//     description: entity created/ modified
//     schema:
//         "$ref": "#/definitions/SearchResult"
//   '403':
//     description: input validation error
//     schema:
////       "$ref": "#/definitions/Response"
//   '500':
//     description: server error
//     schema:
//       "$ref": "#/definitions/Response"
func (c EntityEditController) UpdateEntity() revel.Result {

	type Payload struct {
		Entity models.Entity `json:"entity"`
		Title  string        `json:"title"`
	}
	var (
		err     error
		payload Payload
	)

	err = c.Params.BindJSON(&payload)
	if err != nil {
		log.Println("binding error:", err)
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrorResponse(err, 403))
	}
	go func(passedPayload Payload) {
		log.Println("update entity request", passedPayload.Title)
		existingEntity, err := repositories.EntityRepository{}.GetEntityBy("title", passedPayload.Title)
		if err != nil {
			log.Println("error finding entity:", err)
		}
		passedPayload.Entity.Id = existingEntity.GetId()
		if existingEntity.Title!= passedPayload.Entity.Title{
			titleValue := models.Value{}.
				SetType(ValueType.String).
				SetValueString(passedPayload.Entity.GetTitle()).
				SetDate(time.Now()).
				SetSource("manual edit")
			//TODO: set user as source
			passedPayload.Entity=passedPayload.Entity.SetTitle(titleValue)
		}
		err = repositories.EntityRepository{}.UpdateEntity(passedPayload.Entity)
		if err != nil {
			log.Println("error updating entity:", err)
		}
	}(payload)

	return c.RenderJSON(controllers.BuildSuccessResponse(payload.Entity, 200))
}
