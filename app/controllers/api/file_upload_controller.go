package api

import (
	"GIG-SDK/libraries"
	"GIG-SDK/models"
	"GIG/app/storages"
	"github.com/revel/revel"
	"net/url"
)

type FileUploadController struct {
	*revel.Controller
}

// swagger:operation POST /upload File upload
//
// Upload File
//
// This API allows to upload a file to the server from a remote source url
//
// ---
// produces:
// - application/json
//
// parameters:
//
// - name: upload
//   in: body
//   description: upload object
//   required: true
//   schema:
//       "$ref": "#/definitions/Upload"
//
// security:
//   - Bearer: []
//   - ApiKey: []
//
// responses:
//   '200':
//     description: file uploaded
//     schema:
//         type: string
//   '403':
//     description: input validation error
//     schema:
////       "$ref": "#/definitions/Response"
//   '500':
//     description: server error
//     schema:
//       "$ref": "#/definitions/Response"
func (c FileUploadController) Upload() revel.Result {
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

	tempDir := storages.FileStorageHandler{}.GetCacheDirectory() + upload.GetTitle() + "/"
	tempFile := tempDir + decodedFileName
	if err = libraries.EnsureDirectory(tempDir); err != nil {
		return c.RenderJSON(err)
	}

	if err := libraries.DownloadFile(tempFile, upload.GetSource());
		err != nil {
		return c.RenderJSON(err)
	}

	if err = (storages.FileStorageHandler{}.UploadFile(upload.GetTitle(), tempFile)); err != nil {
		return c.RenderJSON(err)
	}

	c.Response.Status = 200
	return c.RenderJSON("success")
}