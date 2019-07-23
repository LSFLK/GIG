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
func UploadFile(client *minio.Client, bucketName string, filePath string) error {

	if err := client.MakeBucket(bucketName, ""); err != nil {
		// Check to see if we already own this bucket
		exists, errBucketExists := client.BucketExists(bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			fmt.Println("bucket not created")
			return err
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// Upload the file with FPutObject
	if _, err := client.FPutObject(bucketName, utility.ExtractFileName(filePath),
		filePath, minio.PutObjectOptions{ContentType: ""}); err != nil {
		return err
	}

	return nil
}
