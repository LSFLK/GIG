package test_values

import (
	"github.com/lsflk/gig-sdk/enums/ValueType"
	"github.com/lsflk/gig-sdk/models"
	"time"
)

var (
	DateFormat       = "2006-1-2"
	Source0          = "Source0"
	Source           = "Source"
	Source2          = "Source2"
	Source3          = "Source3"
	ValueString      = "~test /tit?le % "
	ValueString2     = "~test /tit?le % 2"
	ValueString3     = "~test /tit?le % 3"
	Date, _          = time.Parse(DateFormat, "2010-5-20")
	Date2, _         = time.Parse(DateFormat, "2010-5-22")
	Date3, _         = time.Parse(DateFormat, "2011-5-22")
	valueType        = ValueType.String
	TestAttributeKey = "test_attribute"

	TestValueObj0 = models.Value{}.
		SetSource(Source0).
		SetValueString(ValueString).
		SetType(valueType)

	TestValueObj = models.Value{}.
		SetSource(Source).
		SetValueString(ValueString).
		SetDate(Date).
		SetType(valueType)

	TestValueObj2 = models.Value{}.
		SetSource(Source2).
		SetValueString(ValueString2).
		SetDate(Date2).
		SetType(valueType)

	TestValueObj3 = models.Value{}.
		SetSource(Source3).
		SetValueString(ValueString3).
		SetDate(Date3).
		SetType(valueType)
)
