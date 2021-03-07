package vangogh_properties

import "github.com/arelate/vangogh_types"

func SupportsProperty(pt vangogh_types.ProductType, property string) bool {
	if !vangogh_types.ValidProductType(pt) {
		return false
	}

	switch property {
	case IdProperty:
		switch pt {
		case vangogh_types.StoreProducts:
			fallthrough
		case vangogh_types.AccountProducts:
			fallthrough
		case vangogh_types.WishlistProducts:
			fallthrough
		case vangogh_types.ApiProductsV1:
			fallthrough
		case vangogh_types.ApiProductsV2:
			return true
		default:
			return false
		}
	case TitleProperty:
		switch pt {
		case vangogh_types.StoreProducts:
			fallthrough
		case vangogh_types.AccountProducts:
			fallthrough
		case vangogh_types.WishlistProducts:
			fallthrough
		case vangogh_types.Details:
			fallthrough
		case vangogh_types.ApiProductsV1:
			fallthrough
		case vangogh_types.ApiProductsV2:
			return true
		default:
			return false
		}
	case DeveloperProperty:
		switch pt {
		case vangogh_types.StoreProducts:
			fallthrough
		case vangogh_types.ApiProductsV2:
			return true
		default:
			return false
		}
	case PublisherProperty:
		switch pt {
		case vangogh_types.StoreProducts:
			fallthrough
		case vangogh_types.ApiProductsV2:
			return true
		default:
			return false
		}
	default:
		return false
	}
}
