package vangogh_properties

import "github.com/arelate/vangogh_types"

var downloadTypeProperties = map[vangogh_types.DownloadType]string{
	vangogh_types.Image:                 ImageProperty,
	vangogh_types.BoxArt:                BoxArtProperty,
	vangogh_types.BackgroundImage:       BackgroundImageProperty,
	vangogh_types.GalaxyBackgroundImage: GalaxyBackgroundImageProperty,
	vangogh_types.Logo:                  LogoProperty,
	vangogh_types.Icon:                  IconProperty,
	vangogh_types.Screenshots:           ScreenshotsProperty,
}

func FromDownloadType(dt vangogh_types.DownloadType) string {
	return downloadTypeProperties[dt]
}
