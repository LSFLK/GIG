package minio

import (
	"GIG/app/utility"
	"fmt"
	"github.com/minio/minio-go"
)

/**
Upload file to minio storage
 */
func UploadFile(client *minio.Client,bucketName string, filePath string) error {

	// Upload the file with FPutObject
	_, err := client.FPutObject(bucketName, utility.ExtractFileName(filePath), filePath, minio.PutObjectOptions{ContentType: ""})
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}
