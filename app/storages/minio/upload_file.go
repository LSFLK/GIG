package minio

import (
	"context"
	"github.com/lsflk/gig-sdk/libraries"
	"github.com/minio/minio-go/v7"
	"log"
)

/*
Upload file to minio storage
*/
func (h Handler) UploadFile(directoryName string, filePath string) error {
	if err := h.Client.MakeBucket(context.Background(), directoryName, minio.MakeBucketOptions{}); err != nil {
		// Check to see if we already own this bucket
		exists, errBucketExists := h.Client.BucketExists(context.Background(), directoryName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", directoryName)
		} else {
			log.Println("bucket not created")
			return err
		}
	} else {
		log.Printf("Successfully created %s\n", directoryName)
	}

	// Upload the file with FPutObject
	if _, err := h.Client.FPutObject(context.Background(), directoryName, libraries.ExtractFileName(filePath),
		filePath, minio.PutObjectOptions{ContentType: ""}); err != nil {
		return err
	}

	return nil
}
