package minio

import (
	"GIG/app/storages/interfaces"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/revel/revel"
	"log"
)

type FileHandler struct {
	interfaces.StorageHandlerInterface
	client         *minio.Client
	cacheDirectory string
}

var minioHandler FileHandler

func (h FileHandler) GetCacheDirectory() string {
	return h.cacheDirectory
}

/*
NewHandler - Always use the NewHandler method to create an instance.
Otherwise, the handler will not be configured
*/
func NewHandler(cacheDirectory string) *FileHandler {
	if minioHandler.client == nil {
		var err error
		endpoint, _ := revel.Config.String("minio.endpoint")
		accessKeyID, _ := revel.Config.String("minio.accessKeyID")
		secretAccessKey, _ := revel.Config.String("minio.secretAccessKey")
		secureUrl, _ := revel.Config.Bool("minio.secureUrl")
		minioHandler.cacheDirectory = cacheDirectory

		// Initialize minio client object.
		minioHandler.client, err = minio.New(endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: secureUrl,
		})
		if err != nil {
			log.Fatal("error connecting to Minio file server")
		}
		log.Println("minio connected")
	}

	return &minioHandler
}
