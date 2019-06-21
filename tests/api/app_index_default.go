package api

import (
	"GIG/app/routes"
)

func (t *TestAPI) TestThatIndexDefaultPageWorks() {
	t.Get(routes.AppController.Index())
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}