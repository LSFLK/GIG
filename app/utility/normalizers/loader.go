package normalizers

import "github.com/revel/revel"

func LoadNormalizers() {
	MapApiUrl, _ = revel.Config.String("normalizer.mapApiUrl")
	MapAppKey, _ = revel.Config.String("normalizer.mapAppKey")

	SearchApiUrl, _ = revel.Config.String("normalizer.searchApiUrl")
	SearchAppKey, _ = revel.Config.String("normalizer.searchAppKey")
	Cx, _ = revel.Config.String("normalizer.cx")
}
