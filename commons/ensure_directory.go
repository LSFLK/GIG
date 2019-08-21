package commons

import (
	"os"
)

func EnsureDirectory(filePath string) error{
	// make directory if not exist
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		createError := os.Mkdir(filePath, os.ModePerm)
		if createError != nil {
			return err
		}
	}
	return nil
}
