package controllers

import (
	"github.com/revel/revel/testing"
)

type TestControllers struct {
	testing.TestSuite
}

func (t *TestControllers) Before() {
	println("Set up")
}

func (t *TestControllers) After() {
	println("Tear down")
}
