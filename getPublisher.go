package vangogh_properties

import (
	"fmt"
	"github.com/arelate/gog_types"
	"github.com/arelate/vangogh_types"
	"github.com/arelate/vangogh_values"
)

func getPublisher(id string, pt vangogh_types.ProductType, mt gog_types.Media) (publisher string, err error) {
	var publisherGetter gog_types.PublisherGetter
	valueReader, err := vangogh_values.NewReader(pt, mt)
	if err != nil {
		return publisher, err
	}

	switch pt {
	case vangogh_types.StoreProducts:
		publisherGetter, err = valueReader.StoreProduct(id)
	case vangogh_types.WishlistProducts:
		publisherGetter, err = valueReader.WishlistProduct(id)
	case vangogh_types.ApiProductsV2:
		publisherGetter, err = valueReader.ApiProductV2(id)
	default:
		return publisher, fmt.Errorf("vangogh_properties: unsupported product type %s", pt)
	}

	if err != nil {
		return publisher, err
	}

	return publisherGetter.GetPublisher(), nil
}
