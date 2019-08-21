package decoders

import (
	"GIG/scriptsorters/etender/model"
	"time"
)

func Decode(result []string) model.ETender {
	sourceDate, _ := time.Parse("01/02/06", result[2])
	closingDate, _ := time.Parse("01/02/06", result[6])
	tender := model.ETender{
		Title:       result[0],
		Company:     result[1],
		SourceDate:  sourceDate,
		Category:    result[3],
		Subcategory: result[4],
		Location:    result[5],
		ClosingDate: closingDate,
		SourceName:  result[7],
		Description: result[8],
		Value:       result[9],
	}

	if tender.Category == tender.Subcategory {
		tender.Subcategory = ""
	}

	return tender
}
