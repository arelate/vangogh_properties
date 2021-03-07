package vangogh_properties

import (
	"fmt"
	"github.com/arelate/gog_types"
	"github.com/arelate/vangogh_types"
	"github.com/arelate/vangogh_values"
)

func GetTitle(id string, pt vangogh_types.ProductType, mt gog_types.Media) (title string, err error) {
	if !SupportsProperty(pt, TitleProperty) {
		return "", fmt.Errorf("vangogh_properties: %s doesn't support property %s", pt, TitleProperty)
	}

	var titleGetter gog_types.TitleGetter
	valueReader, err := vangogh_values.NewReader(pt, mt)
	if err != nil {
		return title, err
	}

	switch pt {
	case vangogh_types.StoreProducts:
		titleGetter, err = valueReader.StoreProduct(id)
	case vangogh_types.AccountProducts:
		titleGetter, err = valueReader.AccountProduct(id)
	case vangogh_types.WishlistProducts:
		titleGetter, err = valueReader.WishlistProduct(id)
	case vangogh_types.Details:
		titleGetter, err = valueReader.Details(id)
	case vangogh_types.ApiProductsV1:
		titleGetter, err = valueReader.ApiProductV1(id)
	case vangogh_types.ApiProductsV2:
		titleGetter, err = valueReader.ApiProductV2(id)
	default:
		return title, fmt.Errorf("unsupported product type %s", pt)
	}

	if err != nil {
		return title, err
	}

	return titleGetter.GetTitle(), nil
}
