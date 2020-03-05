package models

import (
	"time"
)

type SearchResult struct {
	Title      string             `json:"title" bson:"title"`
	Snippet    string             `json:"snippet" bson:"snippet"`
	Categories []string           `json:"categories" bson:"categories"`
	Attributes map[string][]Value `json:"attributes" bson:"attributes"`
	Links      []string           `json:"links" bson:"links"`
	SourceDate time.Time          `json:"source_date" bson:"source_date"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
	ImageURL   string             `json:"image_url" bson:"image_url"`
}

func (s SearchResult) ResultFrom(entity Entity) SearchResult {

	s.Title = entity.GetTitle()
	s.Snippet = entity.GetSnippet()
	s.SourceDate = entity.GetSourceDate()
	s.Categories = entity.GetCategories()
	s.Attributes = entity.GetAttributes()
	s.Links = entity.GetLinks()
	s.CreatedAt = entity.GetCreatedDate()
	s.UpdatedAt = entity.GetUpdatedDate()
	s.ImageURL = entity.GetImageURL()
	return s
}
