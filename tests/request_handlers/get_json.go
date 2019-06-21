package request_handlers

import (
	"GIG/app/utility/request_handlers"
)

func (t *TestRequestHandlers) TestThatGetJsonWorks() {
	link := "https://en.wikipedia.org/w/api.php?action=query&format=json&titles=Sri%20Lanka&prop=extracts&exintro&explaintext"
	result, _ := request_handlers.GetJSON(link)
	t.AssertEqual(len(result),2)
}