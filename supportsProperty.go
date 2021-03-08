package vangogh_properties

import "github.com/arelate/vangogh_types"

// TODO: change all to a map[string]vangogh_types.ProductType

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
	case BoxArtProperty:
		return supportsBoxArtProperty(pt)
	case BackgroundImageProperty:
		return supportsBackgroundImageProperty(pt)
	case GalaxyBackgroundImageProperty:
		return supportsGalaxyBackgroundImageProperty(pt)
	case LogoProperty:
		return supportsLogoProperty(pt)
	case IconProperty:
		return supportsIconProperty(pt)
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

func supportsBoxArtProperty(pt vangogh_types.ProductType) bool {
	switch pt {
	case vangogh_types.ApiProductsV2:
		return true
	default:
		return false
	}
}

func supportsLogoProperty(pt vangogh_types.ProductType) bool {
	switch pt {
	case vangogh_types.ApiProductsV2:
		return true
	default:
		return false
	}
}

func supportsIconProperty(pt vangogh_types.ProductType) bool {
	switch pt {
	case vangogh_types.ApiProductsV1:
		fallthrough
	case vangogh_types.ApiProductsV2:
		return true
	default:
		return false
	}
}

func supportsBackgroundImageProperty(pt vangogh_types.ProductType) bool {
	switch pt {
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

func supportsGalaxyBackgroundImageProperty(pt vangogh_types.ProductType) bool {
	switch pt {
	case vangogh_types.ApiProductsV2:
		return true
	default:
		return false
	}
}
