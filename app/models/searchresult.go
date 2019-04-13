package models

type SearchResult struct {
	Title      string        `json:"title" bson:"title"`
	Content    string        `json:"content" bson:"content"`
}

