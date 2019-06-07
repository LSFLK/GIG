package requesthandlers

import (
	"GIG/app/utility/requesthandlers"
	"github.com/revel/revel/testing"
)

type GetRequestTest struct {
	testing.TestSuite
}

func (t *GetRequestTest) Before() {
	println("Set up")
}

func (t *GetRequestTest) TestThatGetRequestWorks() {
	link := "http://www.buildings.gov.lk/index.php"
	result, _ := requesthandlers.GetRequest(link)
	defer result.Body.Close()
	t.AssertEqual(result.Status,"200 OK")
}

func (t *GetRequestTest) After() {
	println("Tear down")
}
