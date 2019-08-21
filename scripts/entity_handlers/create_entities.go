package entity_handlers

import (
	"GIG/app/models"
	"GIG/scripts"
	"GIG/commons/request_handlers"
	"encoding/json"
)

/**
Create a list of new entities and save to GIG
 */
func CreateEntities(entities []models.Entity) ([]models.Entity, error) {

	resp, err := request_handlers.PostRequest(scripts.ApiUrl+"add-batch", entities)
	if err != nil {
		return entities, err
	}
	json.Unmarshal([]byte(resp), &entities)

	return entities, err
}
