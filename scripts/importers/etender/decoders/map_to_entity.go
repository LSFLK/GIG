package decoders

import (
	"GIG/app/models"
	"GIG/scripts/importers/etender/model"
)

func MapToEntity(tender model.ETender) models.Entity {
	return models.Entity{
	}.
		AddCategory(tender.Category).
		AddCategory(tender.Subcategory).
		SetTitle(models.Value{
			Type:     "string",
			RawValue: tender.Title + " - " + tender.Location,
			Date:     tender.SourceDate,
			Source:   tender.SourceName,
		}).
		SetAttribute("Source Date", models.Value{
			Type:     "date",
			RawValue: tender.SourceDate.String(),
		}).
		SetAttribute("Category", models.Value{
			Type:     "string",
			RawValue: tender.Category,
		}).
		SetAttribute("Subcategory", models.Value{
			Type:     "string",
			RawValue: tender.Subcategory,
		}).
		SetAttribute("Closing Date", models.Value{
			Type:     "date",
			RawValue: tender.ClosingDate.String(),
		}).
		SetAttribute("Source Name", models.Value{
			Type:     "string",
			RawValue: tender.SourceName,
		}).
		SetAttribute("Description", models.Value{
			Type:     "string",
			RawValue: tender.Description,
		}).
		SetAttribute("Value", models.Value{
			Type:     "string",
			RawValue: tender.Value,
		})
}
