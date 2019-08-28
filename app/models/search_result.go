package models

import "encoding/json"

type SearchResult struct {
	Title      string   `json:"title" bson:"title"`
	Snippet    string   `json:"snippet" bson:"snippet"`
	Categories []string `json:"categories" bson:"categories"`
}

func (s SearchResult) ResultFrom(entity Entity) SearchResult {
	jsonAttributes, _ := json.Marshal(entity.Attributes)
	snippet := string(jsonAttributes)
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
	return s
}
