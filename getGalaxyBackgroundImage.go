package vangogh_properties

import (
	"fmt"
	"github.com/arelate/gog_types"
	"github.com/arelate/gog_urls"
	"github.com/arelate/vangogh_types"
	"github.com/arelate/vangogh_values"
)

func getGalaxyBackgroundImage(id string, pt vangogh_types.ProductType, mt gog_types.Media) (galaxyBackgroundImage string, err error) {
	var galaxyBackgroundImageGetter gog_types.GalaxyBackgroundImageGetter
	valueReader, err := vangogh_values.NewReader(pt, mt)
	if err != nil {
		return galaxyBackgroundImage, err
	}

	switch pt {
	case vangogh_types.ApiProductsV2:
		galaxyBackgroundImageGetter, err = valueReader.ApiProductV2(id)
	default:
		return galaxyBackgroundImage, fmt.Errorf("vangogh_properties: unsupported product type %s", pt)
	}

	if err != nil {
		return galaxyBackgroundImage, err
	}

	return gog_urls.ImageId(galaxyBackgroundImageGetter.GetGalaxyBackgroundImage()), nil
}
