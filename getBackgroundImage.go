package vangogh_properties

import (
	"fmt"
	"github.com/arelate/gog_types"
	"github.com/arelate/vangogh_types"
	"github.com/arelate/vangogh_values"
)

func getBackgroundImage(id string, pt vangogh_types.ProductType, mt gog_types.Media) (backgroundImage string, err error) {
	var backgroundImageGetter gog_types.BackgroundImageGetter
	valueReader, err := vangogh_values.NewReader(pt, mt)
	if err != nil {
		return backgroundImage, err
	}

	switch pt {
	case vangogh_types.Details:
		backgroundImageGetter, err = valueReader.Details(id)
	case vangogh_types.ApiProductsV1:
		backgroundImageGetter, err = valueReader.ApiProductV1(id)
	case vangogh_types.ApiProductsV2:
		backgroundImageGetter, err = valueReader.ApiProductV2(id)
	default:
		return backgroundImage, fmt.Errorf("vangogh_properties: unsupported product type %s", pt)
	}

	if err != nil {
		return backgroundImage, err
	}

	return backgroundImageGetter.GetBackgroundImage(), nil
}
