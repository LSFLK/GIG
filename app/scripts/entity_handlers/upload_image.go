package entity_handlers

import (
	"GIG/app/models"
	"GIG/app/scripts"
	"GIG/app/utility/request_handlers"
)

/**
Upload an image through API
 */
func UploadImage(payload models.Upload) error {

	if _, err := request_handlers.PostRequest(scripts.ApiUrl+"upload", payload); err != nil {
		return err
	}
	return nil
}
