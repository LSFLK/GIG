package models

import (
	"GIG/commons"
	"time"
)

type Link struct {
	Title string      `json:"title" bson:"title"`
	Dates []time.Time `json:"dates" bson:"dates"`
}

func (l Link) SetTitle(title string) Link {
	l.Title = title
	return l
}

func (l Link) GetTitle() string {
	return l.Title
}

func (l Link) AddDate(date time.Time) Link {
	if !commons.DateInSlice(l.Dates, date) {
		l.Dates = append(l.Dates, date)
	}
	return l
}

func (l Link) GetDates() []time.Time {
	return l.Dates
}
