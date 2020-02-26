package models

import (
	"strings"
	"time"
)

type Value struct {
	ValueType   string    `json:"value_type" bson:"value_type"`
	ValueString string    `json:"value_string" bson:"value_string"`
	Source      string    `json:"source" bson:"source"`
	Date        time.Time `json:"date" bson:"date"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

func (v Value) SetType(valueType string) Value {
	v.UpdatedAt = time.Now()
	v.ValueType = valueType
	return v
}

func (v Value) GetType() string {
	return v.ValueType
}

func (v Value) SetValueString(value string) Value {
	v.UpdatedAt = time.Now()
	v.ValueString = strings.TrimSpace(value)
	return v
}

func (v Value) GetValueString() string {
	return v.ValueString
}

func (v Value) SetSource(value string) Value {
	v.UpdatedAt = time.Now()
	v.Source = value
	return v
}

func (v Value) GetSource() string {
	return v.Source
}

func (v Value) SetDate(value time.Time) Value {
	v.UpdatedAt = time.Now()
	v.Date = value
	return v
}

func (v Value) GetDate() time.Time {
	return v.Date
}

func (v Value) GetUpdatedDate() time.Time {
	return v.UpdatedAt
}
