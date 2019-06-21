package utility

import (
	"GIG/app/utility"
	"os"
)

func (t *TestUtilities) TestThatDownloadFileWorks() {
	os.Remove("app/cache/download_file_test.ico")
	link := "https://www.wikipedia.org/static/favicon/wikipedia.ico"
	result := utility.DownloadFile("app/cache/download_file_test.ico",link)
	t.AssertEqual(nil, result)
}