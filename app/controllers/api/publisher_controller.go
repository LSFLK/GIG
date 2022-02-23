package api

import (
	"GIG-SDK/models"
	"GIG/app/controllers"
	"GIG/app/publishers/twitter"
	"bufio"
	"bytes"
	"github.com/revel/revel"
	"io"
	"log"
	"os"
)

type PublisherController struct {
	*revel.Controller
}

const (
	chunksize int = 1024
)

func FileToBuffer(data *os.File) (buffer *bytes.Buffer) {

	var (
		part  []byte
		err   error
		count int
	)

	reader := bufio.NewReader(data)
	buffer = bytes.NewBuffer(make([]byte, 0))
	part = make([]byte, chunksize)

	for {
		if count, err = reader.Read(part); err != nil {
			break
		}
		buffer.Write(part[:count])
	}
	if err != io.EOF {
		log.Println("Error Reading ", data.Name(), ": ", err)
	}

	return
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

	mediaId, uploadError := twitter.UploadMedia(entity.ImageURL)
	if uploadError != nil {
		log.Println("media upload error", uploadError)
	}

	publishError := twitter.PublishPost(entity, mediaId)
	if publishError != nil {
		log.Println("post publish error", uploadError)
	}

	return c.RenderJSON(controllers.BuildSuccessResponse("publish request queued.", 200))

}
