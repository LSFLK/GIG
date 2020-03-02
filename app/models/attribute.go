package models

import "sort"

type Attribute struct {
	Name   string  `json:"name" bson:"name"`
	Values []Value `json:"values" bson:"values"`
}

func sortValues(Values []Value) []Value {
	values := Values
	sort.Slice(values, func(i, j int) bool { return values[i].GetDate().Before(values[j].GetDate()) })
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
	a.Values = sortValues(append(a.Values, value))
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

func (a Attribute) GetValues() []Value {
	a.Values = sortValues(a.Values)
	return a.Values
}
