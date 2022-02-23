package api

import (
	"GIG-SDK/models"
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
	log.Println("twitter publish request")
	err = c.Params.BindJSON(&entity)

	if err != nil {
		log.Println("binding error:", err)
		c.Response.Status = 403
		return c.RenderJSON(controllers.BuildErrorResponse(err, 403))
	}

	if err != nil {
		c.Response.Status = 500
		return c.RenderJSON(controllers.BuildErrorResponse(err, 500))
	}

	mediaId, uploadError := twitter_client.UploadMedia(entity.ImageURL)
	if uploadError != nil {
		log.Println("media upload error", uploadError)
	}


	publishError := twitter_client.PublishPost(entity, mediaId)
	if publishError != nil {
		log.Println("post publish error", publishError)
	}

	return c.RenderJSON(controllers.BuildSuccessResponse("publish request queued.", 200))

}
