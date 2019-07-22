package minio

import (
	"GIG/app/utility"
	"fmt"
	"github.com/minio/minio-go"
	"log"
)

/**
Upload file to minio storage
 */
func UploadFile(client *minio.Client,bucketName string, filePath string) error {

	err := client.MakeBucket(bucketName, "")
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := client.BucketExists(bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			return err
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// Upload the file with FPutObject
	_, err = client.FPutObject(bucketName, utility.ExtractFileName(filePath), filePath, minio.PutObjectOptions{ContentType: ""})
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}
