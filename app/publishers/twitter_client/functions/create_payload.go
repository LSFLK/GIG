package functions

import (
	"GIG/app/storages"
	"bytes"
	"io"
	"mime/multipart"
)

func CreatePayload(title string, filename string) (error, *bytes.Buffer, *multipart.Writer) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	localFile, err := storages.FileStorageHandler{}.GetFile(title, filename)
	if err != nil {
		return err, payload, writer
	}
	part1, err := writer.CreateFormFile("media", filename)
	if err != nil {
		return err, payload, writer
	}
	_, err = io.Copy(part1, localFile)
	if err != nil {
		return err, payload, writer
	}
	err = writer.Close()
	if err != nil {
		return err, payload, writer
	}
	return nil, payload, writer
}
