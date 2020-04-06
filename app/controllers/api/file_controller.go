package api

import (
	"GIG-SDK/libraries"
	"GIG-SDK/models"
	"GIG/app/storages"
	"github.com/revel/revel"
	"log"
	"net/url"
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
	c.Response.Status = 400
	if err := c.Params.BindJSON(&upload); err != nil {
		return c.RenderJSON(err)
	}

	decodedFileName, err := url.QueryUnescape(libraries.ExtractFileName(upload.GetSource()))
	if err != nil {
		return c.RenderJSON(err)
	}

	tempDir := storages.FileStorageHandler.GetCacheDirectory() + upload.GetTitle() + "/"
	tempFile := tempDir + decodedFileName
	if err = libraries.EnsureDirectory(tempDir); err != nil {
		return c.RenderJSON(err)
	}

	if err := libraries.DownloadFile(tempFile, upload.GetSource());
		err != nil {
		return c.RenderJSON(err)
	}

	if err = storages.FileStorageHandler.UploadFile(upload.GetTitle(), tempFile); err != nil {
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
	tempDir := storages.FileStorageHandler.GetCacheDirectory() + title + "/"
	sourcePath := tempDir + filename

	if _, err := os.Stat(sourcePath); os.IsNotExist(err) { // if file is not cached
		localFile, err = storages.FileStorageHandler.GetFile(title, filename)
		if err != nil {
			log.Println(err)
			return c.RenderJSON(err)
		}

	} else { // if file is cached
		localFile, err = os.Open(sourcePath)
	}

	c.Response.Status = 200
	return c.RenderFile(localFile, revel.Inline)
}
