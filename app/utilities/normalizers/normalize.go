package normalizers

import (
	"GIG/commons"
	"errors"
	"strings"
)

var StringMinMatchPercentage int

func Normalize(searchString string) (string, error) {

	//using wiki naming registry
	namesArray, err := NormalizeName(searchString)
	if err != nil {
		return "", err
	}

	if len(namesArray) > 0 {
		if commons.StringsMatch(searchString, namesArray[0], StringMinMatchPercentage) {
			return namesArray[0], nil
		}
	}

	// trying to normalize using locations registry
	//locationsArray, err := NormalizeLocation(searchString)
	//if err != nil {
	//	return "", err
	//}
	//if len(locationsArray.Results) > 0 {
	//	return locationsArray.Results[0].FormattedName, err
	//}
	return "", errors.New("no normalizations found")
}

func ProcessNameString(stringValue string) string {
	signature := strings.ToLower(stringValue)
	signature = strings.NewReplacer(
		"%", "",
		"/", "",
		"~", "",
		"?", "",
		"&", "",
		"'", "",
		".", "",
		"(", "",
		")", "",
		"[", "",
		"]", "",
		" etc ", "",
		" etc. ", "",
		" from ", " ",
		",", " ",
		"-", " ",
		" and ", " ",
		" the ", " ",
		" of ", " ",
		" an ", " ",
		" a ", " ",
		" for ", " ",
		" in ", " ",
		" at ", " ",
		" on ", " ",
	).Replace(signature)
	return signature
}
