package entity_handlers

import (
	"GIG/app/models"
	"GIG/scripts"
	"GIG/commons/request_handlers"
	"encoding/json"
)

/**
Create a new entity and save to GIG
 */
func CreateEntity(entity models.Entity) (models.Entity, error) {

	resp, err := request_handlers.PostRequest(scripts.ApiUrl+"add", entity)
	if err != nil {
		return entity, err
	}
	json.Unmarshal([]byte(resp), &entity)
	return entity, err
}
