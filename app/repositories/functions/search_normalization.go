package functions

import (
	"GIG/app/utilities/normalizers"
	"github.com/lsflk/gig-sdk/libraries"
	"github.com/lsflk/gig-sdk/models"
	"log"
)

func SearchNormalizationInCache(normalizedNames []models.NormalizedName, processedEntityTitle string) (isNormalized bool, normalizedTitle string) {
	for _, normalizedName := range normalizedNames {
		if libraries.StringsMatch(processedEntityTitle, normalizedName.GetSearchText(), normalizers.StringMinMatchPercentage) {
			normalizedTitle := normalizedName.GetNormalizedText()
			log.Println("normalization found in cache:", processedEntityTitle, "->", normalizedTitle)
			return true, normalizedTitle
		}
	}
	return false, normalizedTitle
}

func SearchNormalizationInDatabase(normalizedNames []models.Entity, processedEntityTitle string) (isNormalized bool, normalizedTitle string) {
	for _, normalizedName := range normalizedNames {
		if libraries.StringsMatch(processedEntityTitle, libraries.ProcessNameString(normalizedName.GetTitle()), normalizers.StringMinMatchPercentage) {
			normalizedTitle = normalizedName.GetTitle()
			log.Println("normalization found in entity database:", processedEntityTitle, "->", normalizedTitle)
			return true, normalizedTitle
		}
	}
	return false, normalizedTitle
}

func SearchNormalizationInSearchAPI(entityTitle string, processedEntityTitle string) (isNormalized bool, normalizedTitle string) {
	normalizedName, normalizedNameErr := normalizers.Normalize(entityTitle)
	if normalizedNameErr == nil && libraries.StringsMatch(processedEntityTitle, libraries.ProcessNameString(normalizedName), normalizers.StringMinMatchPercentage) {
		log.Println("normalization found in search API:", entityTitle, "->", normalizedTitle)
		//NormalizedNameRepository{}.AddTitleToNormalizationDatabase(entityTitle, normalizedTitle)
		return true, normalizedName
	}
	if normalizedNameErr != nil {
		log.Println("normalization err:", normalizedNameErr)
	}
	return false, normalizedTitle
}

func SearchNormalizationInLocationSearchAPI(entityTitle string) (isNormalized bool, normalizedTitle string) {
	normalizedNameArray, normalizedNameErr := normalizers.NormalizeLocation(entityTitle)
	if len(normalizedNameArray.Results) > 0 {
		normalizedName := normalizedNameArray.Results[0].FormattedName
		if normalizedNameErr == nil {
			log.Println("normalization found in search API:", entityTitle, "->", normalizedTitle)
			//NormalizedNameRepository{}.AddTitleToNormalizationDatabase(entityTitle, normalizedTitle)
			return true, normalizedName
		}
	}
	if normalizedNameErr != nil {
		log.Println("normalization err:", normalizedNameErr)
	}
	return false, normalizedTitle
}
