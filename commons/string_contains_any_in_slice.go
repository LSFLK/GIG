package commons

import "strings"

/**
check if a given string exist in any string in a given string slice
 */
func StringContainsAnyInSlice(slice []string, element string) bool {
	for _, existingElement := range slice {
		if strings.Contains(element,existingElement) {
			return true
		}
	}
	return false
}
