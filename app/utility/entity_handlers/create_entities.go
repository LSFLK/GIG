package entity_handlers

import (
	"GIG/app/models"
	"GIG/app/utility/request_handlers"
	"encoding/json"
)

func CreateEntities(entities []models.Entity) ([]models.Entity, error) {

	resp, err := request_handlers.PostRequest(ApiUrl+"add-batch", entities)
	if err != nil {
		return entities, err
	}
	json.Unmarshal([]byte(resp), &entities)

	return entities, err
}
