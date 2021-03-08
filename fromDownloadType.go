package vangogh_properties

import "github.com/arelate/vangogh_types"

func FromDownloadType(dt vangogh_types.DownloadType) string {
	switch dt {
	case vangogh_types.Image:
		return ImageProperty
	case vangogh_types.BoxArt:
		return BoxArtProperty
	case vangogh_types.Logo:
		return LogoProperty
	case vangogh_types.Icon:
		return IconProperty
	default:
		return ""
	}
}
