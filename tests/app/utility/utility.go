package utility

import (
	"github.com/revel/revel/testing"
)

type TestUtilities struct {
	testing.TestSuite
}

func (t *TestUtilities) Before() {
	println("Set up")
}

func (t *TestUtilities) After() {
	println("Tear down")
}
