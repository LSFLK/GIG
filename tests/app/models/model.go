package models

import (
	"GIG-SDK/models"
	"GIG-SDK/models/ValueType"
	"github.com/revel/revel/testing"
	"time"
)

var (
	source0              = "source0"
	source               = "source"
	source2              = "source2"
	source3              = "source3"
	valueString          = "~test /tit?le % "
	valueString2         = "~test /tit?le % 2"
	valueString3         = "~test /tit?le % 3"
	date, _              = time.Parse("2006-1-2", "2010-5-20")
	date2, _             = time.Parse("2006-1-2", "2010-5-22")
	date25, _             = time.Parse("2006-1-2", "2010-11-22")
	date3, _             = time.Parse("2006-1-2", "2011-5-22")
	valueType            = ValueType.String
	formattedValueString = "2test -title"
	testAttributeKey     = "test_attribute"

	testValueObj0 = models.Value{}.
		SetSource(source0).
		SetValueString(valueString).
		SetType(valueType)

	testValueObj = models.Value{}.
		SetSource(source).
		SetValueString(valueString).
		SetDate(date).
		SetType(valueType)

	testValueObj2 = models.Value{}.
		SetSource(source2).
		SetValueString(valueString2).
		SetDate(date2).
		SetType(valueType)

	testValueObj3 = models.Value{}.
		SetSource(source3).
		SetValueString(valueString3).
		SetDate(date3).
		SetType(valueType)
)

type TestModels struct {
	testing.TestSuite
}

func (t *TestModels) Before() {
	println("Set up")
}

func (t *TestModels) After() {
	println("Tear down")
}
