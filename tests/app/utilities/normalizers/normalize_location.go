package normalizers

import (
	"GIG/app/utilities/normalizers"
)

func (t *TestNormalizers) TestThatNormalizeLocationWorks() {
	result, _ := normalizers.NormalizeLocation("startupx foundry, colombo")
	t.AssertEqual(result.Results[0].FormattedName, "7 Charles Pl, Colombo, Sri Lanka")

}
