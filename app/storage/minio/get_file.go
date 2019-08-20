package minio

import (
	"GIG/app/utility"
	"github.com/minio/minio-go"
	"io"
	"os"
)

/**
Retrieve file from storage
 */
func (h Handler) GetFile(directoryName string, filename string) (*os.File, error) {
	object, err := h.Client.GetObject(directoryName, filename, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer object.Close()
	tempDir := "app/cache/" + directoryName + "/"
	sourcePath := tempDir + filename

	if err = utility.EnsureDirectory(tempDir); err != nil {
		return nil, err
	}

	localFile, err := os.Create(sourcePath)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(localFile, object); err != nil {
		return nil, err
	}
	return localFile, err
}
