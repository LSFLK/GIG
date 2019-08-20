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
func (h Handler)UploadFile(directoryName string, filePath string) error {
	if err := h.Client.MakeBucket(directoryName, ""); err != nil {
		// Check to see if we already own this bucket
		exists, errBucketExists := h.Client.BucketExists(directoryName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", directoryName)
		} else {
			fmt.Println("bucket not created")
			return err
		}
	} else {
		log.Printf("Successfully created %s\n", directoryName)
	}

	// Upload the file with FPutObject
	if _, err := h.Client.FPutObject(directoryName, utility.ExtractFileName(filePath),
		filePath, minio.PutObjectOptions{ContentType: ""}); err != nil {
		return err
	}

	return nil
}
