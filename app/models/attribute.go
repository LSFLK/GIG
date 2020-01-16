package models

import "sort"

type Attribute struct {
	Name   string  `json:"name" bson:"name"`
	Values []Value `json:"values" bson:"values"`
}

/**
Set New Value to Attribute
 */
func (a Attribute) SetValue(value Value) Attribute {
	a.Values = append(a.Values, value)
	return a
}

/**
Get Last Value of Attribute by default
 */
func (a Attribute) GetValue() Value {
	if len(a.Values) == 0 {
		return Value{}
	}
	sort.Slice(a.Values, func(i, j int) bool {return a.Values[i].StartDate.Before(a.Values[j].StartDate)})
	return a.Values[len(a.Values)-1]
}
