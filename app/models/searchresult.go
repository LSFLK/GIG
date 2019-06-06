package models

type SearchResult struct {
	Title   string `json:"title" bson:"title"`
	Snippet []Attribute `json:"snippet" bson:"snippet"`
}