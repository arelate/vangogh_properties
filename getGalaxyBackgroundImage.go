package vangogh_properties

import (
	"fmt"
	"github.com/arelate/gog_types"
	"github.com/arelate/gog_urls"
	"github.com/arelate/vangogh_types"
	"github.com/arelate/vangogh_values"
)

func getGalaxyBackgroundImage(id string, pt vangogh_types.ProductType, mt gog_types.Media) (galaxyBackground string, err error) {
	var galaxyBackgroundGetter gog_types.GalaxyBackgroundGetter
	valueReader, err := vangogh_values.NewReader(pt, mt)
	if err != nil {
		return galaxyBackground, err
	}

	switch pt {
	case vangogh_types.ApiProductsV2:
		galaxyBackgroundGetter, err = valueReader.ApiProductV2(id)
	default:
		return galaxyBackground, fmt.Errorf("vangogh_properties: unsupported product type %s", pt)
	}

	if err != nil {
		return galaxyBackground, err
	}

	return gog_urls.ImageId(galaxyBackgroundGetter.GetGalaxyBackground()), nil
}
