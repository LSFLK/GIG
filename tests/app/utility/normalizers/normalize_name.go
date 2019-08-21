package normalizers

import "GIG/app/utility/normalizers"

func (t *TestNormalizers) TestThatNormalizeNameWorks() {
	result, _ := normalizers.NormalizeName("sri lanka")
	t.AssertEqual(result[0], "Sri Lanka")

}