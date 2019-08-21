package commons

import (
	"github.com/revel/revel/testing"
)

type TestCommons struct {
	testing.TestSuite
}

func (t *TestCommons) Before() {
	println("Set up")
}

func (t *TestCommons) After() {
	println("Tear down")
}
