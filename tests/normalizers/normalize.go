package normalizers

import (
	"GIG/app/utility/normalizers"
)

func (t *TestNormalizers) TestThatNormalizeWorks() {
	result, _ := normalizers.Normalize("sri lanka")
	t.AssertEqual(result, "Sri Lanka")

}