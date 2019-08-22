package api

import (
	"GIG/app/models"
	"GIG/app/storages"
	"GIG/commons"
	"fmt"
	"github.com/revel/revel"
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

	decodedFileName, err := url.QueryUnescape(commons.ExtractFileName(upload.SourceURL))
	if err != nil {
		return c.RenderJSON(err)
	}

	tempFile := storages.FileStorageHandler.GetCacheDirectory() + upload.Title + "/" + decodedFileName
	if err := commons.DownloadFile(tempFile, upload.SourceURL);
		err != nil {
		return c.RenderJSON(err)
	}

	if err = storages.FileStorageHandler.UploadFile(upload.Title, tempFile); err != nil {
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
			fmt.Println(err)
			return c.RenderJSON(err)
		}

	} else { // if file is cached
		localFile, err = os.Open(sourcePath)
	}

	c.Response.Status = 200
	return c.RenderFile(localFile, revel.Inline)
}
