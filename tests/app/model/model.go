package model

import (
	"github.com/revel/revel/testing"
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
