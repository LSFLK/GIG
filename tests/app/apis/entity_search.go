package apis

func (t *TestAPI) TestThatSearchAPIWorks() {
	t.Get("/api/search?query=Sri%20Lanka&limit=10")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}