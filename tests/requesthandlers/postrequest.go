package requesthandlers

import (
	"GIG/app/utility/requesthandlers"
	"github.com/revel/revel/testing"
)

type PostRequestTest struct {
	testing.TestSuite
}

func (t *PostRequestTest) Before() {
	println("Set up")
}

func (t *PostRequestTest) TestThatPostRequestWorks() {
	link := "https://en.wikipedia.org/w/api.php?action=query&format=json&titles=Sri%20Lanka&prop=extracts&exintro&explaintext"
	result, _ := requesthandlers.PostRequest(link,"")
	defer result.Body.Close()
	t.AssertEqual(result.Status,"200 OK")
}

func (t *PostRequestTest) After() {
	println("Tear down")
}
