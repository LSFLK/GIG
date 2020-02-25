package models

import (
	"time"
)

type Value struct {
	valueType string
	rawValue  string
	source    string
	date      time.Time
	updatedAt time.Time
}

func (v *Value) SetType(valueType string) Value {
	v.updatedAt = time.Now()
	v.valueType = valueType
	return *v
}

func (v Value) GetType() string {
	return v.valueType
}

func (v *Value) SetValueString(value string) Value {
	v.updatedAt = time.Now()
	v.rawValue = value
	return *v
}

func (v Value) GetValueString() string {
	return v.rawValue
}

func (v *Value) SetSource(value string) Value {
	v.updatedAt = time.Now()
	v.source = value
	return *v
}

func (v Value) GetSource() string {
	return v.source
}

func (v *Value) SetDate(value time.Time) Value {
	v.updatedAt = time.Now()
	v.date = value
	return *v
}

func (v Value) GetDate() time.Time {
	return v.date
}

func (v Value) GetUpdatedDate() time.Time {
	return v.updatedAt
}
