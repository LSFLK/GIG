package api

import "github.com/revel/revel/testing"

var apiUrl = "http://localhost:9000/api/"

type TestAPI struct {
	testing.TestSuite
}

func (t *TestAPI) Before() {
	println("Set up")
}

func (t *TestAPI) After() {
	println("Tear down")
}