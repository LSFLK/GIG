package utility

import "strings"

/**
check if a given string is contained in any string in a given slice
 */
func StringContainsAnyInSlice(slice []string, element string) bool {
	for _, existingElement := range slice {
		if strings.Contains(element,existingElement) {
			return true
		}
	}
	return false
}
