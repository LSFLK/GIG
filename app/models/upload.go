package models

type Upload struct {
	SourceURL string `json:"source_url" bson:"source_url"`
	Title     string `json:"title" bson:"title"`
}
