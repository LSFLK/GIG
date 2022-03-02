package minio

import (
	"github.com/minio/minio-go"
	"github.com/minio/minio-go/pkg/credentials"
	"github.com/revel/revel"
	"log"
)

type Handler struct {
	Client         *minio.Client
	CacheDirectory string
}

func (h Handler) GetCacheDirectory() string {
	return h.CacheDirectory
}

/**
Always use the NewHandler method to create an instance.
Otherwise the handler will not be configured
 */
func NewHandler(cacheDirectory string) *Handler {
	var err error
	handler := new(Handler)
	endpoint, _ := revel.Config.String("minio.endpoint")
	accessKeyID, _ := revel.Config.String("minio.accessKeyID")
	secretAccessKey, _ := revel.Config.String("minio.secretAccessKey")
	handler.CacheDirectory = cacheDirectory

	// Initialize minio client object.
	handler.Client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Println("error connecting to Minio file server")
		panic(err)
	}
	return handler
}
