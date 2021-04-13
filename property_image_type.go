package vangogh_properties

import "github.com/arelate/vangogh_images"

var imageTypeProperties = map[vangogh_images.ImageType]string{
	vangogh_images.Image:            ImageProperty,
	vangogh_images.BoxArt:           BoxArtProperty,
	vangogh_images.Background:       BackgroundProperty,
	vangogh_images.GalaxyBackground: GalaxyBackgroundProperty,
	vangogh_images.Logo:             LogoProperty,
	vangogh_images.Icon:             IconProperty,
	vangogh_images.Screenshots:      ScreenshotsProperty,
}

func FromImageType(it vangogh_images.ImageType) string {
	return imageTypeProperties[it]
}

func ToImageType(property string) vangogh_images.ImageType {
	for it, prop := range imageTypeProperties {
		if prop == property {
			return it
		}
	}
	return vangogh_images.Unknown
}
