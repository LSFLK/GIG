package normalizers

import (
	"GIG/app/repositories"
	"GIG/app/utilities/normalizers"
)

func (t *TestNormalizers) TestThatNormalizeAPIWorks() {
	result, _ := normalizers.Normalize("sri lanka")
	result2, _ := normalizers.Normalize("ranil")
	result3, _ := normalizers.Normalize("sirisena sri lanka")
	result4, _ := normalizers.Normalize("health ministry sri lanka")
	result5, _ := normalizers.Normalize("election commission sri lanka")
	t.AssertEqual(result, "Sri Lanka")
	t.AssertEqual(result2, "Ranil Wickremesinghe")
	t.AssertEqual(result3, "Maithripala Sirisena")
	t.AssertEqual(result4, "Ministry of Health, Nutrition and Indigenous Medicine")
	t.AssertEqual(result5, "Election Commission of Sri Lanka")

}

func (t *TestNormalizers) TestThatEntityNormalizerWorksWithNormalizerDatabase() {

	result, err := repositories.NormalizeEntityTitle("sri lanka")
	result3, err3 := repositories.NormalizeEntityTitle("All State Bank and their subsidiaries")

	t.AssertEqual(result, "Sri Lanka")
	t.AssertEqual(err, nil)
	t.AssertEqual(result3, "All State Bank and their subsidiaries")
	t.AssertNotEqual(err3, nil)
}
