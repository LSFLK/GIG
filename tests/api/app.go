package api

import (
	"github.com/revel/revel/testing"
	"GIG/app/routes"
)

type AppAPITest struct {
	testing.TestSuite
}

func (t *AppAPITest) Before() {
	println("Set up")
}

func (t *AppAPITest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppAPITest) TestThatIndexDefaultPageWorks() {
	t.Get(routes.AppController.Index())
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppAPITest) After() {
	println("Tear down")
}
