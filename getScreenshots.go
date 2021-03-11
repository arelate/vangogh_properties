package vangogh_properties

import (
	"fmt"
	"github.com/arelate/gog_types"
	"github.com/arelate/vangogh_types"
	"github.com/arelate/vangogh_values"
	"strings"
)

func getScreenshots(id string, pt vangogh_types.ProductType, mt gog_types.Media) (screenshots string, err error) {
	var screenshotsGetter gog_types.ScreenshotsGetter
	valueReader, err := vangogh_values.NewReader(pt, mt)
	if err != nil {
		return screenshots, err
	}

	switch pt {
	case vangogh_types.StoreProducts:
		screenshotsGetter, err = valueReader.StoreProduct(id)
	case vangogh_types.ApiProductsV1:
		screenshotsGetter, err = valueReader.ApiProductV1(id)
	case vangogh_types.ApiProductsV2:
		screenshotsGetter, err = valueReader.ApiProductV2(id)
	default:
		return screenshots, fmt.Errorf("vangogh_properties: unsupported product type %s", pt)
	}

	if err != nil {
		return screenshots, err
	}

	return strings.Join(screenshotsGetter.GetScreenshots(), ","), nil
}
