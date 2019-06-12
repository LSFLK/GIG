package decoders

import (
	"GIG/app/models"
	"GIG/app/utility/importers/etender/model"
)

func MapToEntity(tender model.ETender) models.Entity  {
	return models.Entity{
		Title:    tender.Title + " - " + tender.Location,
	}.
		AddCategory(tender.Category).
		AddCategory(tender.Subcategory).
		AddLink(tender.Company).
		AddLink(tender.Location).
		SetAttribute("Title", models.Value{
			Type:     "string",
			RawValue: tender.Title,
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
		SetAttribute("Location", models.Value{
			Type:     "string",
			RawValue: tender.Location,
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