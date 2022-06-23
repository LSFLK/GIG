package minio

import (
	"context"
	"github.com/lsflk/gig-sdk/libraries"
	"github.com/minio/minio-go/v7"
	"io"
	"os"
)

/*
GetFile - Retrieve file from storage
*/
func (h FileHandler) GetFile(directoryName string, filename string) (*os.File, error) {
	object, err := h.client.GetObject(context.Background(), directoryName, filename, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer object.Close()
	tempDir := h.cacheDirectory + directoryName + "/"
	sourcePath := tempDir + filename

	if err = libraries.EnsureDirectory(tempDir); err != nil {
		return nil, err
	}

	localFile, err := os.Create(sourcePath)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(localFile, object); err != nil {
		os.Remove(sourcePath)
		return nil, err
	}
	return localFile, err
}
