package entity_handlers

import "github.com/revel/revel/testing"

type TestEntityHandlers struct {
	testing.TestSuite
}

func (t *TestEntityHandlers) Before() {
	println("Set up")
}

func (t *TestEntityHandlers) After() {
	println("Tear down")
}