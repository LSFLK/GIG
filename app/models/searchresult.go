package models

type SearchResult struct {
	Title   string `json:"title" bson:"title"`
	Snippet string `json:"snippet" bson:"snippet"`
	Categories []string `json:"categories" bson:"categories"`

}