package interfaces

import "os"

type StorageHandlerInterface interface {
	GetFile(directoryName string, filename string) (*os.File, error)
	UploadFile(directoryName string, filePath string) error
	GetCacheDirectory() string
}
