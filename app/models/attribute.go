package models

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
	return a.Values[len(a.Values)-1]
}
