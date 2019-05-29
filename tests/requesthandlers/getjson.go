package requesthandlers

import (
"GIG/app/utility/requesthandlers"
"github.com/revel/revel/testing"
)

type GetJsonTest struct {
	testing.TestSuite
}

func (t *GetJsonTest) Before() {
	println("Set up")
}

func (t *GetJsonTest) TestThatGetJsonWorks() {
	link := "https://en.wikipedia.org/w/api.php?action=query&format=json&titles=Sri%20Lanka&prop=extracts&exintro&explaintext"
	result, _ := requesthandlers.GetJSON(link)
	t.AssertEqual(len(result),2)
}

func (t *GetJsonTest) After() {
	println("Tear down")
}
