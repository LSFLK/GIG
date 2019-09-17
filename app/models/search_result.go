package models

import (
	"time"
)

type SearchResult struct {
	Title      string    `json:"title" bson:"title"`
	Snippet    string    `json:"snippet" bson:"snippet"`
	Categories []string  `json:"categories" bson:"categories"`
	Links      []string  `json:"links" bson:"links"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
	ImageURL   string    `json:"image_url" bson:"image_url"`
}

func (s SearchResult) ResultFrom(entity Entity) SearchResult {
	snippetAttr, err := entity.GetAttribute("")
	snippet := ""
	if err == nil {
		snippet = snippetAttr.GetValue().RawValue
	}
	snippetAttribute, err := entity.GetAttribute("snippet")
	if err != nil {
		if len(snippet) > 300 {
			snippet = snippet[0:300] + "..."
		}
	} else {
		snippet = snippetAttribute.GetValue().RawValue
	}
	s.Title = entity.Title
	s.Snippet = snippet
	s.Categories = entity.Categories
	s.Links = entity.Links
	s.CreatedAt = entity.CreatedAt
	s.UpdatedAt = entity.UpdatedAt
	s.ImageURL = entity.ImageURL
	return s
}
