package commons

import "gopkg.in/mgo.v2/bson"

/**
check if a given string exists in a given slice
 */
func ObjectIdInSlice(slice []bson.ObjectId, element bson.ObjectId) bool {
	for _, existingElement := range slice {
		if existingElement == element {
			return true
		}
	}
	return false
}
