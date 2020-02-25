package models

import "sort"

type Attribute struct {
	name   string
	values []Value
}

func (a *Attribute) SetName(name string) Attribute {
	a.name = name
	return *a
}

func (a Attribute) GetName() string {
	return a.name
}

/**
Set New Value to Attribute
 */
func (a *Attribute) SetValue(value Value) Attribute {
	a.values = append(a.values, value)
	return *a
}

/**
Get Last Value of Attribute by default
 */
func (a Attribute) GetValue() Value {
	if len(a.values) == 0 {
		return Value{}
	}
	sort.Slice(a.values, func(i, j int) bool { return a.values[i].GetDate().Before(a.values[j].GetDate()) })
	return a.values[len(a.values)-1]
}

func (a Attribute) GetValues() map[string]Value {
	result := make(map[string]Value)
	for _, value := range a.values {
		result[value.GetDate().String()] = value
	}
	return result
}
