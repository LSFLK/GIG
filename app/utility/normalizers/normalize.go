package normalizers

import (
	"GIG/app/utility"
	"GIG/app/utility/normalizers/locations"
	"GIG/app/utility/normalizers/names"
	"errors"
)

func Normalize(searchString string) (string, error) {
	namesArray, err := names.NormalizeName(searchString)
	if err != nil {
		return "", err
	}
	if len(namesArray) > 0 {
		matchPercent := utility.StringMatchPercentage(searchString, namesArray[0])
		if matchPercent > 50 {
			return namesArray[0], nil
		}
	}
	locationsArray, err := locations.NormalizeLocation(searchString)
	if err != nil {
		return "", err
	}
	if len(locationsArray.Results) > 0 {
		return "", errors.New("no normalizations found")
	}
	return locationsArray.Results[0].FormattedName, err
}
