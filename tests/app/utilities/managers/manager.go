package managers

import (
	"GIG-SDK/enums/ValueType"
	"GIG-SDK/models"
	"github.com/revel/revel/testing"
	"time"
)

var (
	source0      = "source0"
	source       = "source"
	source2      = "source2"
	source3      = "source3"
	valueString  = "~test /tit?le % "
	valueString2 = "~test /tit?le % 2"
	valueString3 = "~test /tit?le % 3"
	date, _      = time.Parse("2006-1-2", "2010-5-20")
	date2, _     = time.Parse("2006-1-2", "2010-5-22")
	date3, _     = time.Parse("2006-1-2", "2011-5-22")
	valueType    = ValueType.String

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

type TestManagers struct {
	testing.TestSuite
}

func (t *TestManagers) Before() {
	println("Set up")
}

func (t *TestManagers) After() {
	println("Tear down")
}
