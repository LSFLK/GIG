package api

import (
	"GIG/app/models"
	"GIG/app/storage/minio"
	"GIG/app/utility"
	"fmt"
	"github.com/revel/revel"
	"io"
	"net/url"
	"os"
)

type FileController struct {
	*revel.Controller
}

var baseDir = "app/cache/"

/**
Save file from url
 */
func (c FileController) Upload() revel.Result {
	var (
		upload models.Upload
	)
	c.Response.Status = 400
	if err := c.Params.BindJSON(&upload); err != nil {
		return c.RenderJSON(err)
	}

	decodedFileName, err := url.QueryUnescape(utility.ExtractFileName(upload.SourceURL))
	if err != nil {
		return c.RenderJSON(err)
	}

	tempFile := baseDir + upload.Title + "/" + decodedFileName
	if err := utility.DownloadFile(tempFile, upload.SourceURL);
		err != nil {
		return c.RenderJSON(err)
	}

	minioClient, err := minio.GetClient()
	if err != nil {
		return c.RenderJSON(err)
	}

	if err = minio.UploadFile(minioClient, upload.Title, tempFile); err != nil {
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
	tempDir := baseDir + title + "/"
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

		if err = utility.EnsureDirectory(tempDir); err != nil {
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

	} else { // if file is cached
		localFile, err = os.Open(sourcePath)
	}

	c.Response.Status = 200
	return c.RenderFile(localFile, revel.Inline)
}
