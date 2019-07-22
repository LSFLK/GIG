package minio

import "github.com/minio/minio-go"

/**
Retrieve file from storage
 */
func GetFile(client *minio.Client, bucketName string, name string) (*minio.Object, error) {
	object, err := client.GetObject(bucketName, name, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return object, err
}
