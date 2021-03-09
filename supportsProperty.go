package vangogh_properties

import "github.com/arelate/vangogh_types"

var productTypeProperties = map[vangogh_types.ProductType][]string{
	vangogh_types.StoreProducts:    {IdProperty, TitleProperty, DeveloperProperty, PublisherProperty, ImageProperty},
	vangogh_types.AccountProducts:  {IdProperty, TitleProperty, ImageProperty},
	vangogh_types.WishlistProducts: {IdProperty, TitleProperty, DeveloperProperty, PublisherProperty, ImageProperty},
	vangogh_types.Details:          {TitleProperty, BackgroundImageProperty},
	vangogh_types.ApiProductsV1:    {IdProperty, TitleProperty, IconProperty, BackgroundImageProperty},
	vangogh_types.ApiProductsV2: {
		IdProperty,
		TitleProperty,
		DeveloperProperty,
		PublisherProperty,
		ImageProperty,
		BoxArtProperty,
		LogoProperty,
		IconProperty,
		BackgroundImageProperty,
		GalaxyBackgroundImageProperty,
	},
}

func SupportsProperty(pt vangogh_types.ProductType, property string) bool {
	if !vangogh_types.ValidProductType(pt) {
		return false
	}

	for _, prop := range productTypeProperties[pt] {
		if prop == property {
			return true
		}
	}
	return false
}
