package normalizers

import (
	"GIG/app/models"
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

	entity := repositories.NormalizeEntityTitle(models.Entity{}.SetTitle(models.Value{
		ValueString: "sri lanka",
	}))
	entity1 := repositories.NormalizeEntityTitle(models.Entity{}.SetTitle(models.Value{
		ValueString: "ranil",
	}))
	t.AssertEqual(entity.GetTitle(), "Sri Lanka")
	t.AssertEqual(entity1.GetTitle(), "Ranil Wickremesinghe")
}
