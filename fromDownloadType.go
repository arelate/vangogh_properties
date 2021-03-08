package vangogh_properties

import "github.com/arelate/vangogh_types"

func FromDownloadType(dt vangogh_types.DownloadType) string {
	switch dt {
	case vangogh_types.Image:
		return ImageProperty
	default:
		return ""
	}
}
