package api

import (
	"GIG-SDK/models"
	"GIG/app/constants/error_messages"
	"GIG/app/constants/info_messages"
	"GIG/app/controllers"
	"GIG/app/publishers/twitter_client"
	"github.com/revel/revel"
	"log"
)

type PublisherController struct {
	*revel.Controller
}

func (c PublisherController) Twitter() revel.Result {
	var (
		err    error
		entity models.Entity
	)
	log.Println(info_messages.PublishTwitterRequest)
	err = c.Params.BindJSON(&entity)

	if err != nil {
		log.Println(error_messages.BindingError, err)
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrorResponse(err, 403))
	}

	if err != nil {
		c.Response.Status = 500
		return c.RenderJSON(controllers.BuildErrorResponse(err, 500))
	}

	mediaId, uploadError := twitter_client.UploadMedia(entity.ImageURL)
	if uploadError != nil {
		log.Println(error_messages.MediaUploadError, uploadError)
	}

	publishError := twitter_client.PublishPost(entity, mediaId)
	if publishError != nil {
		log.Println(error_messages.PostPublishError, publishError)
	}

	return c.RenderJSON(controllers.BuildSuccessResponse(info_messages.PublishRequestQueued, 200))

}
