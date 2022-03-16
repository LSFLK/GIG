package managers

import (
	"github.com/revel/revel/testing"
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
