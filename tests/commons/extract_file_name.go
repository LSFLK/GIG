package commons

import (
	"GIG/commons"
)

func (t *TestCommons) TestThatExtractFilenameWorks() {
	filename := "registered%20suppliers%20and%20service%20%20providers%20for%20the%20year%202019%20-%20g1-office%20stationery.pdf"
	link := "/images/pdf/" + filename
	result := commons.ExtractFileName(link)
	t.AssertEqual(filename, result)
}
