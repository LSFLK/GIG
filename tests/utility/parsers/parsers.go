package parsers

import (
	"github.com/revel/revel/testing"
)

type TestParsers struct {
	testing.TestSuite
}

func (t *TestParsers) Before() {
	println("Set up")
}

func (t *TestParsers) After() {
	println("Tear down")
}
