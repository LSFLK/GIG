package normalizers

import (
	"GIG/app/models"
	"GIG/commons"
	"errors"
	"strings"
)

var StringMatchTolerance int

func StringsMatch(string1 string, string2 string) bool {
	matchPercent := commons.StringMatchPercentage(string1, string2)
	return matchPercent > StringMatchTolerance
}

func Normalize(searchString string) (string, error) {

	//using wiki naming registry
	namesArray, err := NormalizeName(searchString)
	if err != nil {
		return "", err
	}
	if len(namesArray) > 0 {
		if StringsMatch(searchString,namesArray[0]) {
			return namesArray[0], nil
		}
	}

	// trying to normalize using locations registry
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
	signature := entity.GetTitle()

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
