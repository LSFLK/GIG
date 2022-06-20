package api

import (
	"GIG/app/storages"
	"github.com/revel/revel"
	"log"
	"os"
)

type FileRetrieveController struct {
	*revel.Controller
}

// swagger:operation GET /images/{title}/{filename}  File retrieve
//
// Retrieve a file from the server
//
// This API allows to retrieve a file from server
//
// ---
// produces:
// - application/json
//
// parameters:
//
// - name: title
//   in: path
//   description: entity title
//   required: true
//   type: string
//
// - name: filename
//   in: path
//   description: filename
//   required: true
//   type: string
//
// responses:
//   '200':
//     description: file
//     schema:
//       type: file
//   '400':
//     description: input parameter validation error
//     schema:
////       "$ref": "#/definitions/Response"
//   '500':
//     description: server error
//     schema:
//       "$ref": "#/definitions/Response"
func (c FileRetrieveController) Retrieve(title string, filename string) revel.Result {
	var (
		localFile *os.File
		err       error
	)

	localFile, err = storages.FileStorageHandler{}.GetFile(title, filename)
	if err != nil {
		log.Println(err)
		c.Response.Status = 400
		return c.RenderJSON("error retrieving file: " + title + "/" + filename)
	}

	c.Response.Status = 200
	return c.RenderFile(localFile, revel.Inline)
}
