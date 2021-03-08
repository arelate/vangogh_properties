package vangogh_properties

import (
	"fmt"
	"github.com/arelate/gog_types"
	"github.com/arelate/vangogh_types"
	"github.com/arelate/vangogh_values"
)

func getLogo(id string, pt vangogh_types.ProductType, mt gog_types.Media) (title string, err error) {
	var logoGetter gog_types.LogoGetter
	valueReader, err := vangogh_values.NewReader(pt, mt)
	if err != nil {
		return title, err
	}

	switch pt {
	case vangogh_types.ApiProductsV2:
		logoGetter, err = valueReader.ApiProductV2(id)
	default:
		return title, fmt.Errorf("vangogh_properties: unsupported product type %s", pt)
	}

	if err != nil {
		return title, err
	}

	return logoGetter.GetLogo(), nil
}
