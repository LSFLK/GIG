package model

import "time"

type ETender struct {
	Title       string    `json:"title"`
	Company     string    `json:"company"`
	SourceDate  time.Time `json:"source_date"`
	Category    string    `json:"category"`
	Subcategory string    `json:"subcategory"`
	Location    string    `json:"subcategory"`
	ClosingDate time.Time `json:"closing_date"`
	SourceName  string    `json:"source_name"`
	Description string    `json:"description"`
	Value       string    `json:"value"`
}
