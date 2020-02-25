package models

import "sort"

type Attribute struct {
	Name   string  `json:"name" bson:"name"`
	Values []Value `json:"values" bson:"values"`
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
	sort.Slice(a.Values, func(i, j int) bool { return a.Values[i].GetDate().Before(a.Values[j].GetDate()) })
	return a.Values[len(a.Values)-1]
}

func (a Attribute) GetValues() map[string]Value {
	result := make(map[string]Value)
	for _, value := range a.Values {
		result[value.GetDate().String()] = value
	}
	return result
}
