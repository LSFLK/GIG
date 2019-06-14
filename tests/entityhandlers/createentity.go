package entityhandlers

import (
	"GIG/app/models"
	"GIG/app/utility/requesthandlers"
	"encoding/json"
	"io/ioutil"
)

func CreateEntity(entity models.Entity) (models.Entity, error) {

	resp, saveErr := requesthandlers.PostRequest(ApiUrl, entity)
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
