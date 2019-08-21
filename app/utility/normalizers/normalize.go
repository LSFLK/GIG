package normalizers

import (
	"GIG/commons"
	"errors"
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
	locationsArray, err := NormalizeLocation(searchString)
	if err != nil {
		return "", err
	}
	if len(locationsArray.Results) == 0 {
		return "", errors.New("no normalizations found")
	}
	return locationsArray.Results[0].FormattedName, err
}
