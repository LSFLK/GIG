package commons

import (
	"GIG/commons"
	"os"
)

func (t *TestCommons) TestThatDownloadFileWorks() {
	os.Remove("app/cache/test.pdf")
	link := "https://s1.q4cdn.com/806093406/files/doc_downloads/test.pdf"
	result := commons.DownloadFile("app/cache/test.pdf",link)
	t.AssertEqual(nil, result)
}