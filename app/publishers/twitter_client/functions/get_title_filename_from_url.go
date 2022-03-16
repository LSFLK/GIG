package functions

import (
	"github.com/tomogoma/go-typed-errors"
	"strings"
)

func GetTitleAndFilenameFromUrl(url string) (err error, title string, filename string) {
	imageUrl := strings.Split(url, "/")
	if len(imageUrl) == 3 { // if a valid image exist
		return nil, imageUrl[1], imageUrl[2]
	}
	return errors.New("invalid url"), "", ""
}
