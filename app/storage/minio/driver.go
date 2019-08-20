package minio

import (
	"github.com/minio/minio-go"
	"github.com/revel/revel"
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
func NewHandler() *Handler {
	var err error
	handler := new(Handler)
	endpoint, _ := revel.Config.String("minio.endpoint")
	accessKeyID, _ := revel.Config.String("minio.accessKeyID")
	secretAccessKey, _ := revel.Config.String("minio.secretAccessKey")
	handler.CacheDirectory, _ = revel.Config.String("minio.cache")

	// Initialize minio client object.
	handler.Client, err = minio.New(endpoint, accessKeyID, secretAccessKey, false)
	if err != nil {
		panic(err)
	}
	return handler
}
