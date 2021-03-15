package vangogh_properties

import "github.com/arelate/vangogh_types"

var imageTypeProperties = map[vangogh_types.ImageType]string{
	vangogh_types.Image:            ImageProperty,
	vangogh_types.BoxArt:           BoxArtProperty,
	vangogh_types.Background:       BackgroundProperty,
	vangogh_types.GalaxyBackground: GalaxyBackgroundProperty,
	vangogh_types.Logo:             LogoProperty,
	vangogh_types.Icon:             IconProperty,
	vangogh_types.Screenshots:      ScreenshotsProperty,
}

func FromImageType(it vangogh_types.ImageType) string {
	return imageTypeProperties[it]
}

func ToImageType(property string) vangogh_types.ImageType {
	for it, prop := range imageTypeProperties {
		if prop == property {
			return it
		}
	}
	return vangogh_types.UnknownImageType
}
