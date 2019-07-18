package api

import (
	"GIG/app/storage/minio"
	"fmt"
	"github.com/revel/revel"
)

type ImageController struct {
	*revel.Controller
}

/**
Save file from url
 */
func (c ImageController) Upload() revel.Result {
	// Check to see if we already own this bucket (which happens if you run this twice)
	minioClient, err := minio.GetClient()
	if err != nil {
		fmt.Println(err)
		c.Response.Status = 400
		return c.RenderJSON(err)
	}

	err = minio.UploadFile(minioClient, "gig", "/home/umayanga/Downloads/routes.csv")
	if err != nil {
		c.Response.Status = 400
		return c.RenderJSON(err)
	}
	c.Response.Status = 200
	return c.RenderJSON("success")
}
