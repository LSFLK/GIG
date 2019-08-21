package commons

import (
	"GIG/commons"
	"os"
)

func (t *TestCommons) TestThatDownloadFileWorks() {
	os.Remove("app/cache/download_file_test.ico")
	link := "https://www.wikipedia.org/static/favicon/wikipedia.ico"
	result := commons.DownloadFile("app/cache/download_file_test.ico",link)
	t.AssertEqual(nil, result)
}