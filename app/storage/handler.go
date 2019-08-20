package storage

import (
	"GIG/app/storage/minio"
	"os"
)

var FileStorageHandler IHandler

type IHandler interface {
	GetFile(directoryName string, filename string) (*os.File, error)
	UploadFile(directoryName string, filePath string) error
	GetCacheDirectory() string
}

func LoadStorageHandler() {
	FileStorageHandler = minio.NewHandler()
}
