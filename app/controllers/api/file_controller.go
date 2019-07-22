package api

import (
	"GIG/app/models"
	"GIG/app/storage/minio"
	"GIG/app/utility"
	"fmt"
	"github.com/revel/revel"
	"io"
	"os"
)

type FileController struct {
	*revel.Controller
}

/**
Save file from url
 */
func (c FileController) Upload() revel.Result {
	var (
		upload models.Upload
	)
	err := c.Params.BindJSON(&upload)

	if err != nil {
		fmt.Println(err)
		return c.RenderJSON(err)
	}
	c.Response.Status = 400

	minioClient, err := minio.GetClient()
	if err != nil {
		fmt.Println(err)
		return c.RenderJSON(err)
	}

	err = minio.UploadFile(minioClient, upload.Title, upload.SourceURL)
	if err != nil {
		return c.RenderJSON(err)
	}
	c.Response.Status = 200
	return c.RenderJSON("success")
}

/**
Retrieve file from storage
 */
func (c FileController) Retrieve(title string, filename string) revel.Result {
	c.Response.Status = 400
	var localFile *os.File
	tempDir := "app/cache/" + title + "/"
	sourcePath := tempDir + filename

	if _, err := os.Stat(sourcePath); os.IsNotExist(err) { // if file is not cached
		minioClient, err := minio.GetClient()
		if err != nil {
			fmt.Println(err)
			return c.RenderJSON(err)
		}

		file, err := minio.GetFile(minioClient, title, filename)
		defer file.Close()
		if err != nil {
			return c.RenderJSON(err)
		}

		err = utility.EnsureDirectory(tempDir)
		if err != nil {
			fmt.Println(err)
			return c.RenderJSON(err)
		}
		localFile, err := os.Create(sourcePath)
		if err != nil {
			return c.RenderJSON(err)
		}
		if _, err = io.Copy(localFile, file); err != nil {
			return c.RenderJSON(err)
		}
		c.Response.Status = 200
		return c.RenderFile(localFile, revel.Inline)
	} else {	// if file is cached
		localFile, err = os.Open(sourcePath)
	}

	c.Response.Status = 200
	return c.RenderFile(localFile, revel.Inline)
}
