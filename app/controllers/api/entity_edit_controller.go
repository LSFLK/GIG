package api

import (
	"GIG/app/constants/error_messages"
	"GIG/app/constants/info_messages"
	"GIG/app/controllers"
	"GIG/app/repositories"
	"GIG/app/services/authentication"
	"GIG/app/services/entity_operations"
	"errors"
	"github.com/lsflk/gig-sdk/enums/ValueType"
	"github.com/lsflk/gig-sdk/models"
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
	log.Println(info_messages.EntityCreateBatch)
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
				_, err := repositories.EntityRepository{}.AddEntity(entity)
				if err != nil {
					log.Println(error_messages.EntityCreateError, e)
				}
			}(e)

		}

		wg.Wait()
	}(entitiesList)

	c.Response.Status = 200
	return c.RenderJSON(controllers.BuildSuccessResponse(info_messages.EntityCreateBatchQueued, 200))
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
	log.Println(info_messages.EntityCreate)
	err = c.Params.BindJSON(&entity)
	if err != nil {
		log.Println(error_messages.BindingError, err)
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrorResponse(err, 403))
	}

	go func(newEntity models.Entity) {
		entity, err = repositories.EntityRepository{}.AddEntity(newEntity)
		if err != nil {
			log.Println(error_messages.EntityCreationError, err)
		}
	}(entity)

	return c.RenderJSON(controllers.BuildSuccessResponse(info_messages.EntityCreateQueued, 200))

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
		entity models.Entity
		err    error
	)
	log.Println(info_messages.EntityTerminate)
	err = c.Params.BindJSON(&entity)
	if err != nil {
		log.Println(error_messages.BindingError, err)
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrorResponse(err, 403))
	}

	if entity.GetTitle() == "" && len(entity.GetCategories()) == 0 {
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrorResponse(errors.New(error_messages.TitleCategoryRequired), 400))
	}

	if entity.GetSourceDate().IsZero() || entity.GetSource() == "" {
		c.Response.Status = 400
		return c.RenderJSON(controllers.BuildErrorResponse(errors.New(error_messages.TerminationDateSourceRequired), 400))
	}

	go entity_operations.HandleEntityTermination(entity)

	return c.RenderJSON(controllers.BuildSuccessResponse(info_messages.EntityTerminateQueued, 200))
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
		log.Println(error_messages.BindingError, err)
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrorResponse(err, 403))
	}
	log.Println(info_messages.EntityDelete, entity.Title)

	go func(entityToDelete models.Entity) {
		existingEntity, err := repositories.EntityRepository{}.GetEntityByPreviousTitle(entityToDelete.Title, time.Now())
		if err != nil {
			log.Println(error_messages.EntityFindError, err)
		}

		err = repositories.EntityRepository{}.DeleteEntity(existingEntity)
		if err != nil {
			log.Println(error_messages.EntityDeleteError, err)
		}
	}(entity)

	return c.RenderJSON(controllers.BuildSuccessResponse(info_messages.EntityDeleteQueued, 200))
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
		log.Println(error_messages.BindingError, err)
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrorResponse(err, 403))
	}
	go func(passedPayload Payload) {
		log.Println(info_messages.EntityUpdate, passedPayload.Title)
		existingEntity, err := repositories.EntityRepository{}.GetEntityBy("title", passedPayload.Title)
		if err != nil {
			log.Println(error_messages.EntityFindError, err)
		}
		passedPayload.Entity.Id = existingEntity.GetId()

		user, _, err := authentication.GetAuthUser(c.Request.Header)
		if err != nil {
			log.Println("trying to get authenticated user error: ", err)
		}
		if existingEntity.Title != passedPayload.Entity.Title {
			titleValue := models.Value{}.
				SetType(ValueType.String).
				SetValueString(passedPayload.Entity.GetTitle()).
				SetDate(time.Now()).
				SetSource(user.Email)
			passedPayload.Entity = passedPayload.Entity.SetTitle(titleValue)
		}
		err = repositories.EntityRepository{}.UpdateEntity(passedPayload.Entity)
		if err != nil {
			log.Println(error_messages.EntityUpdateError, err)
		}
	}(payload)

	return c.RenderJSON(controllers.BuildSuccessResponse(payload.Entity, 200))
}
