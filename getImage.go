package vangogh_properties

import (
	"fmt"
	"github.com/arelate/gog_types"
	"github.com/arelate/vangogh_types"
	"github.com/arelate/vangogh_values"
)

func getImage(id string, pt vangogh_types.ProductType, mt gog_types.Media) (title string, err error) {
	var imageGetter gog_types.ImageGetter
	valueReader, err := vangogh_values.NewReader(pt, mt)
	if err != nil {
		return title, err
	}

	switch pt {
	case vangogh_types.StoreProducts:
		imageGetter, err = valueReader.StoreProduct(id)
	case vangogh_types.AccountProducts:
		imageGetter, err = valueReader.AccountProduct(id)
	case vangogh_types.WishlistProducts:
		imageGetter, err = valueReader.WishlistProduct(id)
	case vangogh_types.ApiProductsV2:
		imageGetter, err = valueReader.ApiProductV2(id)
	default:
		return title, fmt.Errorf("unsupported product type %s", pt)
	}

	if err != nil {
		return title, err
	}

	return imageGetter.GetImage(), nil
}
