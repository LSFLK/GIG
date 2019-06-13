package utility

import "strings"

/**
extract the main domain from a given source path
 */
func ExtractDomain(link string) string {
	splitUrl := strings.Split(link, "/")
	return splitUrl[2]
}