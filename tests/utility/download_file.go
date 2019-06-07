package utility

import (
	"GIG/app/utility"
	"github.com/revel/revel/testing"
	"os"
)

type DownloadFileTest struct {
	testing.TestSuite
}

func (t *DownloadFileTest) Before() {
	println("Set up")
}

func (t *DownloadFileTest) TestThatDownloadFileWorks() {
	os.Remove("app/cache/downloadfiletest.ico")
	link := "https://www.wikipedia.org/static/favicon/wikipedia.ico"
	result := utility.DownloadFile("app/cache/downloadfiletest.ico",link)
	t.AssertEqual(nil, result)
}

func (t *DownloadFileTest) After() {
	println("Tear down")
}
