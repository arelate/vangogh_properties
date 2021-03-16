package vangogh_properties

import (
	"github.com/arelate/vangogh_types"
)

var supportedProperties = map[vangogh_types.ProductType][]string{
	vangogh_types.AccountProducts: {IdProperty, TitleProperty, ImageProperty},
	vangogh_types.ApiProductsV1:   {IdProperty, TitleProperty, IconProperty, BackgroundProperty, ScreenshotsProperty},
	vangogh_types.ApiProductsV2:   {IdProperty, TitleProperty, DeveloperProperty, PublisherProperty, ImageProperty, BoxArtProperty, IconProperty, LogoProperty, BackgroundProperty, GalaxyBackgroundProperty, ScreenshotsProperty},
	vangogh_types.Details:         {TitleProperty, BackgroundProperty},
	vangogh_types.StoreProducts:   {IdProperty, TitleProperty, DeveloperProperty, PublisherProperty, ImageProperty, ScreenshotsProperty},
}

func SupportsProperty(pt vangogh_types.ProductType, property string) bool {
	for _, supportedProperty := range supportedProperties[pt] {
		if property == supportedProperty {
			return true
		}
	}
	return false
}
