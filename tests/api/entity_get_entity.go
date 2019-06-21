package api

func (t *TestAPI) TestThatGetEntityAPIWorks() {
	t.Get("/api/get/Sri%20Lanka")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}