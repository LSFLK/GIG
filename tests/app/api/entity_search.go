package api

func (t *TestAPI) TestThatSearchAPIWorks() {
	t.Get("/api/search?query=Sri%20Lanka")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}