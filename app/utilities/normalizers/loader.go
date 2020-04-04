package normalizers

import (
	"github.com/revel/revel"
)

func LoadNormalizers() {
	MapApiUrl, _ = revel.Config.String("normalizer.mapApiUrl")
	MapAppKey, _ = revel.Config.String("normalizer.mapAppKey")
	StringMinMatchPercentage, _ = revel.Config.Int("normalizer.minMatchPercentage")
}
