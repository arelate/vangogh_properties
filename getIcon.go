package vangogh_properties

import (
	"fmt"
	"github.com/arelate/gog_types"
	"github.com/arelate/gog_urls"
	"github.com/arelate/vangogh_types"
	"github.com/arelate/vangogh_values"
)

func getIcon(id string, pt vangogh_types.ProductType, mt gog_types.Media) (icon string, err error) {
	var iconGetter gog_types.IconGetter
	valueReader, err := vangogh_values.NewReader(pt, mt)
	if err != nil {
		return icon, err
	}

	switch pt {
	case vangogh_types.ApiProductsV1:
		iconGetter, err = valueReader.ApiProductV1(id)
	case vangogh_types.ApiProductsV2:
		iconGetter, err = valueReader.ApiProductV2(id)
	default:
		return icon, fmt.Errorf("vangogh_properties: unsupported product type %s", pt)
	}

	if err != nil {
		return icon, err
	}

	return gog_urls.ImageId(iconGetter.GetIcon()), nil
}
