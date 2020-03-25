package models

import (
	"sort"
	"time"
)

type Attribute struct {
	Name   string  `json:"name" bson:"name"`
	Values []Value `json:"values" bson:"values"`
}

func sortValues(Values []Value, order string) []Value {
	values := Values
	sort.Slice(values, func(i, j int) bool {
		if order == "DESC" {
			return values[i].GetDate().After(values[j].GetDate())
		} else {
			return values[i].GetDate().Before(values[j].GetDate())
		}
	})
	return values
}

func (a Attribute) SetName(name string) Attribute {
	a.Name = name
	return a
}

func (a Attribute) GetName() string {
	return a.Name
}

/**
Set New Value to Attribute
 */
func (a Attribute) SetValue(value Value) Attribute {
	a.Values = sortValues(append(a.Values, value), "ASC")
	return a
}

/**
Get Last Value of Attribute by default
 */
func (a Attribute) GetValue() Value {
	if len(a.Values) == 0 {
		return Value{}
	}
	return a.GetValues()[len(a.Values)-1]
}

/**
Get Value of Attribute by date
 */
func (a Attribute) GetValueByDate(date time.Time) Value {
	sortedValues := sortValues(a.Values, "DESC")

	// pick the value with highest date lower than or equal to the given date
	for _, value := range sortedValues {

		if !value.GetDate().After(date) {
			return value
		}
	}
	//return the raw
	return Value{}
}

func (a Attribute) GetValues() []Value {
	a.Values = sortValues(a.Values, "ASC")
	return a.Values
}
