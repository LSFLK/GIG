package commons

import "strings"

/**
extract filename from a given source path
 */
func ExtractFileName(link string) string {
	splitUrl := strings.Split(link, "/")
	return splitUrl[len(splitUrl)-1]
}