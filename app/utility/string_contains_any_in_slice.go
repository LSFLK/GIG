package utility

import "strings"

/**
return true if a given string exists in a given slice
 */
func StringContainsAnyInSlice(slice []string, element string) bool {
	for _, existingElement := range slice {
		if strings.Contains(element,existingElement) {
			return true
		}
	}
	return false
}
