package storages

import (
	"GIG-SDK/libraries"
	"GIG/app/storages/minio"
	"github.com/revel/revel"
	"log"
	"os"
)

var FileStorageHandler IHandler

type IHandler interface {
	GetFile(directoryName string, filename string) (*os.File, error)
	UploadFile(directoryName string, filePath string) error
	GetCacheDirectory() string
}

func LoadStorageHandler() {
	cacheDirectory, _ := revel.Config.String("file.cache")

	if cacheDirectory == "" { // default value
		cacheDirectory = "app/cache"
	}

	if err := libraries.EnsureDirectory(cacheDirectory); err != nil {
		log.Fatal(err)
	}

	FileStorageHandler = minio.NewHandler(cacheDirectory) //change storage handler here
}
