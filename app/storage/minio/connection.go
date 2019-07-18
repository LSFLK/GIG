package minio

import (
	"github.com/minio/minio-go"
	"github.com/revel/revel"
	"log"
)

func GetClient()(*minio.Client,error) {
	endpoint,_ := revel.Config.String("minio.endpoint")
	accessKeyID,_ := revel.Config.String("minio.accessKeyID")
	secretAccessKey,_ := revel.Config.String("minio.secretAccessKey")

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, false)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return minioClient,err
}
