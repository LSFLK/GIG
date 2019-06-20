package entity_handlers

import (
	"GIG/app/models"
	"GIG/app/utility/request_handlers"
	"encoding/json"
	"io/ioutil"
)

func CreateEntity(entity models.Entity) (models.Entity, error) {

	resp, saveErr := request_handlers.PostRequest(ApiUrl+"add", entity)
	if saveErr != nil {
		return entity, saveErr
	}
	respBody, bodyError := ioutil.ReadAll(resp.Body)
	if bodyError != nil {
		return entity, bodyError
	}
	json.Unmarshal(respBody, &entity)
	resp.Body.Close()
	return entity, bodyError
}

func CreateEntities(entities []models.Entity) ([]models.Entity, error) {

	resp, saveErr := request_handlers.PostRequest(ApiUrl+"add-batch", entities)
	if saveErr != nil {
		return entities, saveErr
	}
	respBody, bodyError := ioutil.ReadAll(resp.Body)
	if bodyError != nil {
		return entities, bodyError
	}
	json.Unmarshal(respBody, &entities)
	resp.Body.Close()

	return entities, bodyError
}
