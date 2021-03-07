package vangogh_properties

import (
	"fmt"
	"github.com/arelate/gog_types"
	"github.com/arelate/vangogh_types"
	"github.com/arelate/vangogh_values"
)

func getDeveloper(id string, pt vangogh_types.ProductType, mt gog_types.Media) (developer string, err error) {
	var developerGetter gog_types.DeveloperGetter
	valueReader, err := vangogh_values.NewReader(pt, mt)
	if err != nil {
		return developer, err
	}

	switch pt {
	case vangogh_types.StoreProducts:
		developerGetter, err = valueReader.StoreProduct(id)
	case vangogh_types.ApiProductsV2:
		developerGetter, err = valueReader.ApiProductV2(id)
	default:
		return developer, fmt.Errorf("unsupported product type %s", pt)
	}

	if err != nil {
		return developer, err
	}

	return developerGetter.GetDeveloper(), nil
}
