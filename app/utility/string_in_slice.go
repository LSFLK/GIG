package utility

/**
check if a given string exists in a given slice
 */
func StringInSlice(slice []string, element string) bool {
	for _, existingElement := range slice {
		if existingElement == element {
			return true
		}
	}
	return false
}
