package repositories

import (
	"github.com/revel/revel/testing"
)

type TestRepositories struct {
	testing.TestSuite
}

func (t *TestRepositories) Before() {
	println("Set up")
}

func (t *TestRepositories) After() {
	println("Tear down")
}
