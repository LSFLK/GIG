package api

import (
	"GIG/scripts"
	"github.com/revel/revel/testing"
)

var apiUrl = scripts.ApiUrl

type TestAPI struct {
	testing.TestSuite
}

func (t *TestAPI) Before() {
	println("Set up")
}

func (t *TestAPI) After() {
	println("Tear down")
}