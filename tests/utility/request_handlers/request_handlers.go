package request_handlers

import (
	"github.com/revel/revel/testing"
)

type TestRequestHandlers struct {
	testing.TestSuite
}

func (t *TestRequestHandlers) Before() {
	println("Set up")
}

func (t *TestRequestHandlers) After() {
	println("Tear down")
}

