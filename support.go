package vangogh_properties

import (
	"github.com/arelate/vangogh_products"
)

var supportedProperties = map[vangogh_products.ProductType][]string{
	vangogh_products.AccountProducts: {
		IdProperty,
		TitleProperty,
		ImageProperty,
		RatingProperty,
		OperatingSystemsProperty,
		SlugProperty,
	},
	vangogh_products.ApiProductsV1: {
		IdProperty,
		TitleProperty,
		IconProperty,
		BackgroundProperty,
		ScreenshotsProperty,
		VideoIdProperty,
		OperatingSystemsProperty,
		SlugProperty,
		GOGReleaseDate,
	},
	vangogh_products.ApiProductsV2: {
		IdProperty,
		TitleProperty,
		DevelopersProperty,
		PublisherProperty,
		ImageProperty,
		BoxArtProperty,
		IconProperty,
		LogoProperty,
		BackgroundProperty,
		GalaxyBackgroundProperty,
		ScreenshotsProperty,
		IncludesGamesProperty,
		IsIncludedByGamesProperty,
		RequiresGamesProperty,
		IsRequiredByGamesProperty,
		GenresProperty,
		FeaturesProperty,
		SeriesProperty,
		VideoIdProperty,
		OperatingSystemsProperty,
		LanguageCodeProperty,
		GlobalReleaseDate,
		GOGReleaseDate,
	},
	vangogh_products.Details: {
		TitleProperty,
		BackgroundProperty,
		FeaturesProperty,
		TagIdProperty,
		GOGReleaseDate,
	},
	vangogh_products.StoreProducts: {
		IdProperty,
		TitleProperty,
		DevelopersProperty,
		PublisherProperty,
		ImageProperty,
		ScreenshotsProperty,
		RatingProperty,
		GenresProperty,
		VideoIdProperty,
		OperatingSystemsProperty,
		SlugProperty,
		GlobalReleaseDate,
		GOGReleaseDate,
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

func Supported(pt vangogh_products.ProductType, properties []string) []string {
	supported := make([]string, 0, len(properties))
	for _, prop := range properties {
		if SupportsProperty(pt, prop) {
			supported = append(supported, prop)
		}
	}
	return supported
}
