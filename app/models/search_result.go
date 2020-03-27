package models

import (
	"time"
)

type SearchResult struct {
	Title      string               `json:"title" bson:"title"`
	Snippet    string               `json:"snippet" bson:"snippet"`
	Categories []string             `json:"categories" bson:"categories"`
	Attributes map[string]Attribute `json:"attributes" bson:"attributes"`
	Links      []Link               `json:"links" bson:"links"`
	SourceDate time.Time            `json:"source_date" bson:"source_date"`
	CreatedAt  time.Time            `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time            `json:"updated_at" bson:"updated_at"`
	ImageURL   string               `json:"image_url" bson:"image_url"`
}

func (s SearchResult) ResultFrom(entity Entity, attributes []string) SearchResult {

	s.Title = entity.GetTitle()
	s.Snippet = entity.GetSnippet()
	s.SourceDate = entity.GetSourceDate()
	s.Categories = entity.GetCategories()
	s.Links = entity.GetLinks()
	s.CreatedAt = entity.GetCreatedDate()
	s.UpdatedAt = entity.GetUpdatedDate()
	s.ImageURL = entity.GetImageURL()

	if len(attributes) == 0 {
		s.Attributes = entity.GetAttributes()
	} else {
		s.Attributes = make(map[string]Attribute)
		for _, attribute := range attributes {
			existingAttribute, attributeErr := entity.GetAttribute(attribute)
			if attributeErr == nil {
				s.Attributes[attribute] = existingAttribute
			}
		}
	}

	return s
}
