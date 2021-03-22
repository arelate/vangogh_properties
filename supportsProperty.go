package vangogh_properties

import (
	"github.com/arelate/vangogh_products"
)

var supportedProperties = map[vangogh_products.ProductType][]string{
	vangogh_products.AccountProducts: {
		IdProperty,
		TitleProperty,
		ImageProperty,
	},
	vangogh_products.ApiProductsV1: {
		IdProperty,
		TitleProperty,
		IconProperty,
		BackgroundProperty,
		ScreenshotsProperty,
	},
	vangogh_products.ApiProductsV2: {
		IdProperty,
		TitleProperty,
		DeveloperProperty,
		PublisherProperty,
		ImageProperty,
		BoxArtProperty,
		IconProperty,
		LogoProperty,
		BackgroundProperty,
		GalaxyBackgroundProperty,
		ScreenshotsProperty,
	},
	vangogh_products.Details: {
		TitleProperty,
		BackgroundProperty,
	},
	vangogh_products.StoreProducts: {
		IdProperty,
		TitleProperty,
		DeveloperProperty,
		PublisherProperty,
		ImageProperty,
		ScreenshotsProperty,
	},
}

func SupportsProperty(pt vangogh_products.ProductType, property string) bool {
	for _, supportedProperty := range supportedProperties[pt] {
		if property == supportedProperty {
			return true
		}
	}
	return false
}
