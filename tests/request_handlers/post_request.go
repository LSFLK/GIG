package request_handlers

import (
	"GIG/app/utility/request_handlers"
)

func (t *TestRequestHandlers) TestThatPostRequestWorks() {
	link := "https://en.wikipedia.org/w/api.php?action=query&format=json&titles=Sri%20Lanka&prop=extracts&exintro&explaintext"
	result, _ := request_handlers.PostRequest(link,"")
	defer result.Body.Close()
	t.AssertEqual(result.Status,"200 OK")
}