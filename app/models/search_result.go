package models

import "encoding/json"

type SearchResult struct {
	Title      string   `json:"title" bson:"title"`
	Snippet    string   `json:"snippet" bson:"snippet"`
	Categories []string `json:"categories" bson:"categories"`
}

func (s SearchResult) ResultFrom(entity Entity) SearchResult {
	jsonAttributes, _ := json.Marshal(entity.Attributes)
	stringAttributes := string(jsonAttributes)
	if len(stringAttributes) > 300 {
		stringAttributes = stringAttributes[0:300] + "..."
	}
	s.Title = entity.Title
	s.Snippet = stringAttributes
	s.Categories = entity.Categories
	return s
}
