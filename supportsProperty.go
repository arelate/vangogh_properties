package vangogh_properties

import "github.com/arelate/vangogh_types"

func SupportsProperty(pt vangogh_types.ProductType, property string) bool {
	if !vangogh_types.ValidProductType(pt) {
		return false
	}

	switch property {
	case IdProperty:
		return supportsIdProperty(pt)
	case TitleProperty:
		return supportsTitleProperty(pt)
	case DeveloperProperty:
		return supportsDeveloperProperty(pt)
	case PublisherProperty:
		return supportsPublisherProperty(pt)
	case ImageProperty:
		return supportsImageProperty(pt)
	default:
		return false
	}
}

func supportsIdProperty(pt vangogh_types.ProductType) bool {
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
}

func supportsTitleProperty(pt vangogh_types.ProductType) bool {
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
}

func supportsDeveloperProperty(pt vangogh_types.ProductType) bool {
	switch pt {
	case vangogh_types.StoreProducts:
		fallthrough
	case vangogh_types.WishlistProducts:
		fallthrough
	case vangogh_types.ApiProductsV2:
		return true
	default:
		return false
	}
}

func supportsPublisherProperty(pt vangogh_types.ProductType) bool {
	switch pt {
	case vangogh_types.StoreProducts:
		fallthrough
	case vangogh_types.WishlistProducts:
		fallthrough
	case vangogh_types.ApiProductsV2:
		return true
	default:
		return false
	}
}

func supportsImageProperty(pt vangogh_types.ProductType) bool {
	switch pt {
	case vangogh_types.StoreProducts:
		fallthrough
	case vangogh_types.AccountProducts:
		fallthrough
	case vangogh_types.WishlistProducts:
		fallthrough
	case vangogh_types.ApiProductsV2:
		return true
	default:
		return false
	}
}
