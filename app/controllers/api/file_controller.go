package api

import (
	"GIG/app/models"
	"GIG/app/storage"
	"GIG/app/utility"
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

	decodedFileName, err := url.QueryUnescape(utility.ExtractFileName(upload.SourceURL))
	if err != nil {
		return c.RenderJSON(err)
	}

	tempFile := storage.FileStorageHandler.GetCacheDirectory() + upload.Title + "/" + decodedFileName
	if err := utility.DownloadFile(tempFile, upload.SourceURL);
		err != nil {
		return c.RenderJSON(err)
	}

	if err = storage.FileStorageHandler.UploadFile(upload.Title, tempFile); err != nil {
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
	tempDir := storage.FileStorageHandler.GetCacheDirectory() + title + "/"
	sourcePath := tempDir + filename

	if _, err := os.Stat(sourcePath); os.IsNotExist(err) { // if file is not cached
		if err != nil {
			fmt.Println(err)
			return c.RenderJSON(err)
		}
		localFile, err = storage.FileStorageHandler.GetFile(title, filename)
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
