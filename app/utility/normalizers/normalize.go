package normalizers

import (
	"GIG/app/models"
	"GIG/app/repository"
	"GIG/app/utility"
	"errors"
)

/*
Normalize function should not be called directly from non server functions
such as utilities since the function depends on server's database configuration for caching;
instead use the normalizer APIs.
 */
func Normalize(searchString string) (string, error) {
	// check if the searchString has already being normalized
	normalizedName, err := repository.GetNomralizedNameBy("searchText", searchString)
	if err == nil {
		return normalizedName.NormalizedText, nil
	}
	namesArray, err := NormalizeName(searchString)
	if err != nil {
		return "", err
	}
	if len(namesArray) > 0 {
		matchPercent := utility.StringMatchPercentage(searchString, namesArray[0])
		if matchPercent > 50 {
			repository.AddNormalizedName(models.NormalizedName{SearchText: searchString, NormalizedText: namesArray[0]})
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
	repository.AddNormalizedName(models.NormalizedName{SearchText: searchString, NormalizedText: locationsArray.Results[0].FormattedName})
	return locationsArray.Results[0].FormattedName, err
}
