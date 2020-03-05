package normalizers

import (
	"GIG/app/models"
	"GIG/commons"
	"errors"
	"strings"
)

var StringMatchTolerance int

func Normalize(searchString string) (string, error) {
	namesArray, err := NormalizeName(searchString)
	if err != nil {
		return "", err
	}
	if len(namesArray) > 0 {
		matchPercent := commons.StringMatchPercentage(searchString, namesArray[0])
		if matchPercent > StringMatchTolerance {
			return namesArray[0], nil
		}
	}
	/**
	TODO: given a entity title
	1. derive a unique signature from the title
	2. check if the signature already exist in the database. if exist merge entities
	3. if does not exist, try to find a normalized name for the title. If found title, merge entities
			check for close signature matches and return if a valid match is found
				log the matches in a separate table
			if a valid match is not found create the entity with the existing name
				log the titles of failed normalizations

	 */
	locationsArray, err := NormalizeLocation(searchString)
	if err != nil {
		return "", err
	}
	if len(locationsArray.Results) == 0 {
		return "", errors.New("no normalizations found")
	}
	return locationsArray.Results[0].FormattedName, err
}

func GenerateEntitySignature(entity models.Entity) string {
	signature:=entity.GetTitle()

	signature = strings.NewReplacer(
		"%", "",
		"/", "",
		"~", "",
		"?", "",
		"&", "",
		"'", "",
		".", "",
		",", " ",
		"-", " ",
		" and ", " ",
		" the ", " ",
		" of ", " ",
		" an ", " ",
		" a ", " ",
		" a ", " ",
	).Replace(signature)
	return signature
}
